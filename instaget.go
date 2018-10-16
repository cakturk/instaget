package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	queryHash = "5b0222df65d7f6659c9b82246780caa7"
)

type graphImageParser struct {
	json map[string]interface{}
}

func (g *graphImageParser) displayResources() ([]*displayResource, error) {
	var res []*displayResource
	dec := rawDecoder{
		dic: g.json,
	}
	dec.keyToArray("display_resources")
	for i := range dec.slice {
		dec.arrayNthElemToMap(i)
		src := dec.dic["src"].(string)
		width := int(dec.dic["config_width"].(float64))
		height := int(dec.dic["config_height"].(float64))
		res = append(res, &displayResource{
			src:    src,
			width:  width,
			height: height,
		})
	}
	return res, dec.err
}

type displayResource struct {
	src           string
	width, height int
}

func (d *displayResource) String() string {
	return fmt.Sprintf("src: %s s%dx%d", d.src, d.width, d.height)
}

type graphSidecarParser struct {
	json map[string]interface{}
}

func (g *graphSidecarParser) sidecarEdges() ([]*graphImageParser, error) {
	var s []*graphImageParser
	dec := rawDecoder{
		dic: g.json,
	}
	dec.keyToMap("edge_sidecar_to_children")
	dec.keyToArray("edges")
	for i := range dec.slice {
		edgeDec := dec
		edgeDec.arrayNthElemToMap(i)
		edgeDec.keyToMap("node")
		if edgeDec.err != nil {
			return nil, edgeDec.err
		}
		s = append(s, &graphImageParser{
			json: edgeDec.dic,
		})
	}
	return s, dec.err
}

type rawDecoder struct {
	dic   map[string]interface{}
	slice []interface{}
	err   error
}

func (r *rawDecoder) reset() {
	r.err = nil
}

func (r *rawDecoder) keyToMap(key string) {
	if r.err != nil {
		return
	}
	m, ok := r.dic[key]
	if !ok {
		r.err = fmt.Errorf("key: '%s' is not found", key)
		return
	}
	newm, ok := m.(map[string]interface{})
	if !ok {
		r.err = fmt.Errorf("could not convert map[%s] to map", key)
		return
	}
	r.dic = newm
}

func (r *rawDecoder) keyToArray(key string) {
	if r.err != nil {
		return
	}
	m, ok := r.dic[key]
	if !ok {
		r.err = fmt.Errorf("key: '%s' is not found", key)
		return
	}
	newsl, ok := m.([]interface{})
	if !ok {
		r.err = fmt.Errorf("could not convert map[%s] to []interface{}", key)
		return
	}
	r.slice = newsl
}

func (r *rawDecoder) arrayNthElemToMap(i int) {
	if r.err != nil {
		return
	}
	if i >= len(r.slice) {
		r.err = fmt.Errorf("array index: %d out of bounds", i)
		return
	}
	m, ok := r.slice[i].(map[string]interface{})
	if !ok {
		r.err = fmt.Errorf("could not convert arr[%d] to map", i)
		return
	}
	r.dic = m
}

type pageType int

const (
	unknownType pageType = iota
	graphImage
	graphSidecar
	profilePage
)

func getType(data string) (pageType, map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		return unknownType, nil, err
	}
	dec := rawDecoder{dic: m}
	dec.keyToMap("entry_data")
	dec.keyToArray("PostPage")
	if dec.err == nil {
		dec.arrayNthElemToMap(0)
		dec.keyToMap("graphql")
		dec.keyToMap("shortcode_media")
		if dec.err != nil {
			return unknownType, nil, err
		}
		typ, ok := dec.dic["__typename"]
		if !ok {
			return unknownType, nil, fmt.Errorf("key: '__typename' is not found")
		}
		switch typ {
		case "GraphImage":
			return graphImage, dec.dic, nil
		case "GraphSidecar":
			return graphSidecar, dec.dic, nil
		}
		return unknownType, nil, fmt.Errorf("unrecognized __typename")
	}
	dec.reset()
	dec.keyToArray("ProfilePage")
	if dec.err == nil {
		dec.arrayNthElemToMap(0)
		dec.keyToMap("graphql")
		if dec.err != nil {
			return unknownType, nil, err
		}
		return profilePage, dec.dic, nil
	}
	return unknownType, nil, fmt.Errorf("unrecognized JSON structure")
}

func getJSON(r io.Reader) (string, error) {
	n, err := findJSONNode(r)
	if err != nil {
		return "", err
	}
	return extractJSONString(n)
}

func extractJSONString(n *html.Node) (string, error) {
	data := n.Data[:len(n.Data)-1]
	idx := strings.Index(n.Data, "{")
	if idx == -1 {
		return data, fmt.Errorf("malformed JSON data")
	}
	return data[idx:], nil
}

func findJSONNode(r io.Reader) (*html.Node, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var prevNode, snode *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			prevNode = n
		case html.TextNode:
			if prevNode.Data != "script" {
				return
			}
			for _, a := range prevNode.Attr {
				if a.Key != "type" {
					continue
				}
				if !strings.HasPrefix(n.Data, "window._sharedData") {
					continue
				}
				snode = n
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if snode == nil {
		return nil, fmt.Errorf("missing script element")
	}
	return snode, nil
}

func httpGet(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("instaget.httpGet: %s", resp.Status)
	}
	return resp, nil
}

func downloadFile(urlStr string) error {
	resp, err := httpGet(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	out, err := os.Create(path.Base(u.Path))
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func scrapeImages(u string) error {
	resp, err := httpGet(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := getJSON(resp.Body)
	if err != nil {
		return err
	}
	typ, jdata, err := getType(f)
	if err != nil {
		return err
	}

	switch typ {
	case graphImage:
		p := graphImageParser{json: jdata}
		r, err := p.displayResources()
		if err != nil {
			return err
		}
		return downloadFile(r[len(r)-1].src)
	case graphSidecar:
		p := graphSidecarParser{json: jdata}
		s, err := p.sidecarEdges()
		if err != nil {
			return err
		}
		var wg sync.WaitGroup
		for _, g := range s {
			res, err := g.displayResources()
			if err != nil {
				return err
			}
			wg.Add(1)
			go func(dr *displayResource) {
				downloadFile(dr.src)
				wg.Done()
			}(res[len(res)-1])
		}
		wg.Wait()
	case profilePage:
		fmt.Println("profile page is not yet supported")
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("requires a URL")
		os.Exit(1)
	}
	if _, err := url.Parse(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := scrapeImages(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
