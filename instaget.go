package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

//                             shortcode
// https://www.instagram.com/p/BpAGN0pBrSH/?__a=1

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	queryURL  = "https://www.instagram.com/graphql/query"
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
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("instaget.httpGet: %s", resp.Status)
	}
	return resp, nil
}

func xhr(u *url.URL, query url.Values, rhxGis string) (resp *http.Response, err error) {
	u.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	signature := md5.Sum([]byte(fmt.Sprintf("%s:%s", rhxGis, query.Get("variables"))))
	// let the golang handle content-encoding automagically by not setting
	// an accept-encoding header.
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("x-instagram-gis", hex.EncodeToString(signature[:]))
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("instaget.xhr: %s", resp.Status)
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

func downloadResources(done chan struct{}, errc chan<- error, urls <-chan string) {
	for u := range urls {
		if err := downloadFile(u); err != nil {
			select {
			case errc <- err:
			case <-done:
				break
			}
		}
		select {
		case <-done:
			break
		default:
		}
	}
}

func doShortcodeRequest(shortcode string) (*ShortcodeQueryResponse, error) {
	resp, err := httpGet("https://www.instagram.com/p/" + shortcode + "/?__a=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	p := new(ShortcodeQueryResponse)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type paginationQuery struct {
	ID    string `json:"id"`
	First int    `json:"first"`
	After string `json:"after"`
}

type pager interface {
	id() string
	endCursor() string
	rhxGis() string
}

func getNextPage(id string, endCursor string, rhxGis string) (*PaginationQueryResponse, error) {
	query := url.Values{}
	query.Add("query_hash", queryHash)
	pr := &paginationQuery{
		ID:    id,
		First: 12,
		After: endCursor,
	}
	b, err := json.Marshal(pr)
	if err != nil {
		return nil, err
	}
	query.Add("variables", string(b))
	u, err := url.Parse(queryURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	resp, err := xhr(u, query, rhxGis)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	qresp := new(PaginationQueryResponse)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(qresp)
	if err != nil {
		return nil, err
	}
	return qresp, nil
}

type urlLister interface {
	listURLs() []string
}

func scrapeProfilePage(paths chan<- string, p *ProfilePostPage) error {
	tm := &p.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia
	// Handle the first page first as a special case, and then do
	// pagination requests.
	for i := range tm.Edges {
		n := &tm.Edges[i].Node
		switch {
		case n.Typename == "GraphImage":
			paths <- n.DisplayURL
		case n.Typename == "GraphSidecar":
			resp, err := doShortcodeRequest(n.Shortcode)
			if err != nil {
				return err
			}
			edges := resp.Graphql.ShortcodeMedia.EdgeSidecarToChildren.Edges
			for i := range edges {
				paths <- edges[i].Node.DisplayResources[2].Src
			}
		case n.Typename == "GraphVideo":
			resp, err := doShortcodeRequest(n.Shortcode)
			if err != nil {
				return err
			}
			paths <- resp.Graphql.ShortcodeMedia.VideoURL
		default:
		}
	}
	var hasNext bool
	hasNext = tm.PageInfo.HasNextPage
	user := &p.EntryData.ProfilePage[0].Graphql.User
	endCursor := user.EdgeOwnerToTimelineMedia.PageInfo.EndCursor
	rhxGis := p.RhxGis
	count := 0
	for hasNext {
		resp, err := getNextPage(user.ID, endCursor, rhxGis)
		if err != nil {
			return err
		}
		for _, u := range resp.listURLs() {
			paths <- u
		}
		hasNext = resp.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
		endCursor = resp.Data.User.EdgeOwnerToTimelineMedia.PageInfo.EndCursor
		count++
		if count > 3 {
			break
		}
	}
	return nil
}

func scrapePostPage(paths chan<- string, p *ProfilePostPage) error {
	for _, u := range p.listURLs() {
		paths <- u
	}
	return nil
}

func scrapeImages(u string) error {
	resp, err := httpGet(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s, err := getJSON(resp.Body)
	if err != nil {
		return err
	}
	p := new(ProfilePostPage)
	err = json.Unmarshal([]byte(s), p)
	if err != nil {
		return err
	}
	paths := make(chan string)
	errc := make(chan error, 1)
	done := make(chan struct{})
	defer close(done)
	const numWorkers = 10
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			downloadResources(done, errc, paths)
			wg.Done()
		}()
	}
	switch {
	case len(p.EntryData.ProfilePage) > 0:
		err = scrapeProfilePage(paths, p)
	case len(p.EntryData.PostPage) > 0:
		err = scrapePostPage(paths, p)
	default:
		err = errors.New("instaget.scrapeImages: unrecognized page type")
	}
	close(paths)
	wg.Wait()
	close(errc)
	if err != nil {
		return err
	}
	if err = <-errc; err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "requires a URL")
		os.Exit(1)
	}
	if _, err := url.Parse(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	jar, _ := cookiejar.New(nil)
	http.DefaultClient.Jar = jar
	if err := scrapeImages(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
