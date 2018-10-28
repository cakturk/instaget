package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	queryURL  = "https://www.instagram.com/graphql/query"
	queryHash = "5b0222df65d7f6659c9b82246780caa7"
)

func extractJSON(r io.Reader) (string, error) {
	n, err := findJSONNode(r)
	if err != nil {
		return "", err
	}
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

type timeSource interface {
	time() time.Time
}

type rangeStatus int

const (
	cont rangeStatus = iota + 1
	inRange
	outOfRange
)

type rangeInfo interface {
	includes(b timeSource) rangeStatus
}

func scrapeProfilePage(ri rangeInfo, paths chan<- string, p *ProfilePostPage) error {
	tm := &p.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia
	// Handle the first page first as a special case, and then do
	// pagination requests.
	for i := range tm.Edges {
		n := &tm.Edges[i].Node
		switch ri.includes(n) {
		case cont:
			continue
		case outOfRange:
			return nil
		case inRange:
		}
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
	for hasNext {
		resp, err := getNextPage(user.ID, endCursor, rhxGis)
		if err != nil {
			return err
		}
		urls, keepGoing := resp.listURLs(ri)
		for _, u := range urls {
			paths <- u
		}
		if !keepGoing {
			return nil
		}
		hasNext = resp.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
		endCursor = resp.Data.User.EdgeOwnerToTimelineMedia.PageInfo.EndCursor
	}
	return nil
}

func scrapePostPage(paths chan<- string, p *ProfilePostPage) error {
	for _, u := range p.listURLs() {
		paths <- u
	}
	return nil
}

func scrapeImages(ri rangeInfo, u string) error {
	resp, err := httpGet(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s, err := extractJSON(resp.Body)
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
		err = scrapeProfilePage(ri, paths, p)
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

var layouts = [...]string{
	"2006-01-02 15:04",
	"2006-01-02",
}

func parseTime(val string) (time.Time, error) {
	var err error
	for _, l := range layouts {
		var tim time.Time
		tim, err = time.ParseInLocation(l, val, time.Local)
		if err == nil {
			return tim, nil
		}
	}
	return time.Time{}, err
}

func createRangeInfo() (rangeInfo, error) {
	if *from != "" && *offset != -1 {
		return nil, errors.New("mutual exclusive options 'offset' and 'from'")
	}
	if *to != "" && *count > 0 {
		fmt.Println(*to, *count)
		return nil, errors.New("mutual exclusive options 'count' and 'to'")
	}
	const (
		flagFrom  = 0x01
		flagTo    = 0x02
		flagOff   = 0x04
		flagCount = 0x08

		rTimeRange      = flagFrom | flagTo
		rCountRange     = flagOff | flagCount
		rCountRange2    = flagCount
		rCountTimeRange = flagOff | flagTo
		rTimeCountRange = flagFrom | flagCount
	)
	var (
		err      error
		flags    uint
		timeFrom time.Time
		timeTo   time.Time
	)
	if *from != "" {
		timeFrom, err = parseTime(*from)
		if err != nil {
			return nil, err
		}
		flags |= flagFrom
	}
	if *to != "" {
		timeTo, err = parseTime(*to)
		if err != nil {
			return nil, err
		}
		flags |= flagTo
	}
	if *offset > -1 {
		flags |= flagOff
	}
	if *count > 0 {
		flags |= flagCount
	}
	switch flags {
	case rTimeRange:
		return &timeRange{start: timeFrom, end: timeTo}, nil
	case rCountRange, rCountRange2:
		off := 0
		if *offset > -1 {
			off = *offset
		}
		return &countRange{off: off, count: *count}, nil
	case rCountTimeRange:
		return &countTimeRange{off: *offset, to: timeTo}, nil
	case rTimeCountRange:
		return &timeCountRange{from: timeFrom, count: *count}, nil
	}
	return nopRange{}, nil
}

type timeCountRange struct {
	from  time.Time
	count int
	curr  int
}

func (t *timeCountRange) includes(i timeSource) rangeStatus {
	tim := i.time()
	if tim.After(t.from) {
		return cont
	}
	if t.curr < t.count {
		t.curr++
		return inRange
	}
	return outOfRange
}

type countTimeRange struct {
	off  int
	curr int
	to   time.Time
}

func (t *countTimeRange) includes(i timeSource) rangeStatus {
	tim := i.time()
	if tim.Before(t.to) || tim.Equal(t.to) {
		return outOfRange
	}
	if tim.After(t.to) {
		defer func() { t.curr++ }()
		if t.curr < t.off {
			return cont
		}
		return inRange
	}
	return outOfRange
}

type countRange struct {
	off   int
	count int
	next  int
}

func (c *countRange) includes(timeSource) rangeStatus {
	if c.next < c.off {
		c.next++
		return cont
	}
	if c.next >= c.off && c.next < c.off+c.count {
		c.next++
		return inRange
	}
	return outOfRange
}

type timeRange struct {
	start time.Time
	end   time.Time
}

func (t *timeRange) includes(i timeSource) rangeStatus {
	c := i.time()
	if c.After(t.start) {
		return cont
	}
	if c.Equal(t.start) {
		return inRange
	}
	if c.Before(t.start) && c.After(t.end) {
		return inRange
	}
	return outOfRange
}

type nopRange struct{}

func (nopRange) includes(timeSource) rangeStatus {
	return inRange
}

var (
	// from and to are in reverse-chronological order
	from   = flag.String("from", "", "Download posts on or older than the specified date")
	to     = flag.String("to", "", "Download posts to or newer than the specified date")
	offset = flag.Int("offset", -1, "Starting post")
	count  = flag.Int("count", 0, "Downloads up to count posts at offset 'offset' (from the start of the timeline)")
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "requires a URL")
		os.Exit(1)
	}
	rngInfo, err := createRangeInfo()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := url.Parse(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	jar, _ := cookiejar.New(nil)
	http.DefaultClient.Jar = jar
	if err := scrapeImages(rngInfo, flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
