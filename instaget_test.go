package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func readAllOrDie(s string, t *testing.T) []byte {
	buf, err := ioutil.ReadFile(s)
	if err != nil {
		t.Fatal(err)
	}
	return buf
}

func TestGetType(t *testing.T) {
	check := func(e error) {
		if e != nil {
			t.Fatal(e)
		}
	}
	cases := []struct {
		path string
		want pageType
	}{
		{"json/instagram-single-pic.json", graphImage},
		{"json/instagram-multi-pic.json", graphSidecar},
		{"json/instagram-profile-page.json", profilePage},
	}
	for _, c := range cases {
		b, err := ioutil.ReadFile(c.path)
		check(err)
		typ, _, err := getType(string(b))
		check(err)
		if typ != c.want {
			t.Errorf("got: %d, want: %d", typ, c.want)
		}
	}
}

func TestGraphImageParser(t *testing.T) {
	buf := readAllOrDie("json/instagram-single-pic.json", t)
	typ, data, _ := getType(string(buf))
	if typ != graphImage {
		t.Errorf("got: %d, want: %d", typ, graphImage)
	}
	p := graphImageParser{
		json: data,
	}
	want := []*displayResource{
		{"https://instagram.fist1-2.fna.fbcdn.net/vp/d86df4991804aa99a1977435c88fe336/5C5048D4/t51.2885-15/sh0.08/e35/s640x640/41753641_400121580525237_8881904291674358247_n.jpg", 640, 640},
		{"https://instagram.fist1-2.fna.fbcdn.net/vp/3b536d0a10a3df76aa7f1d0af902356e/5C4383D4/t51.2885-15/sh0.08/e35/s750x750/41753641_400121580525237_8881904291674358247_n.jpg", 750, 750},
		{"https://instagram.fist1-2.fna.fbcdn.net/vp/c2672b2daac89f3bf5c668ecd8a096d3/5C4F0431/t51.2885-15/e35/41753641_400121580525237_8881904291674358247_n.jpg", 1080, 1080},
	}
	got, _ := p.displayResources()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func TestGraphSidecarParser(t *testing.T) {
	buf := readAllOrDie("json/instagram-multi-pic.json", t)
	typ, data, _ := getType(string(buf))
	if typ != graphSidecar {
		t.Errorf("got: %d, want: %d", typ, graphSidecar)
	}
	p := graphSidecarParser{
		json: data,
	}
	cases := []struct {
		res []*displayResource
	}{
		// edge node 1
		{
			[]*displayResource{
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/2c55d12e0a63ac6a4e75aa0c0035e273/5C442671/t51.2885-15/sh0.08/e35/s640x640/40758827_2138611023072230_4073975203662780931_n.jpg", 640, 640},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/9579ca6bad768cff2d858ea27d260b81/5C3EC3B5/t51.2885-15/sh0.08/e35/s750x750/40758827_2138611023072230_4073975203662780931_n.jpg", 750, 750},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/bc6cc41a5373dc92b32782f009a524e1/5C4445CB/t51.2885-15/e35/40758827_2138611023072230_4073975203662780931_n.jpg", 1080, 1080},
			},
		},
		// edge node 2
		{

			[]*displayResource{
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/9a0db8f1bba3978babfdb35cd19baa5f/5C6251E5/t51.2885-15/sh0.08/e35/s640x640/41441981_319071475313278_5721220910286835828_n.jpg", 640, 640},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/b607cb95c1384655386508a6c1782bc6/5C4941E5/t51.2885-15/sh0.08/e35/s750x750/41441981_319071475313278_5721220910286835828_n.jpg", 750, 750},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/2e913c6c72b4608dfbcaddf65f25deaa/5C5D3E00/t51.2885-15/e35/41441981_319071475313278_5721220910286835828_n.jpg", 1080, 1080},
			},
		},
		// edge node 3
		{

			[]*displayResource{
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/23635d191fdaf4730e4752d67eed9607/5C4EC496/t51.2885-15/sh0.08/e35/s640x640/41073830_1443782439056724_8345211079608769101_n.jpg", 640, 640},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/97d05499d5acde628bb861f5c4ea2775/5C596152/t51.2885-15/sh0.08/e35/s750x750/41073830_1443782439056724_8345211079608769101_n.jpg", 750, 750},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/72b25a95f5a73fb62ab06e744f4c202f/5C47D92C/t51.2885-15/e35/41073830_1443782439056724_8345211079608769101_n.jpg", 1080, 1080},
			},
		},
		// edge node 4
		{

			[]*displayResource{
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/3319d49441016ccae90f0fc95e8a08ef/5C6217B2/t51.2885-15/sh0.08/e35/s640x640/41438271_297182174203680_1114131638392175833_n.jpg", 640, 640},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/4dd37d9cdccc60b4a31989108e225a27/5C4A07B2/t51.2885-15/sh0.08/e35/s750x750/41438271_297182174203680_1114131638392175833_n.jpg", 750, 750},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/2765f1330b541a3bcf946f3821da9377/5C609457/t51.2885-15/e35/41438271_297182174203680_1114131638392175833_n.jpg", 1080, 1080},
			},
		},
		// edge node 5
		{

			[]*displayResource{
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/ba74822e33a8192c3af011c82bd854fa/5C3F7944/t51.2885-15/sh0.08/e35/s640x640/41013725_540667323027121_5550989628702074250_n.jpg", 640, 640},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/d2f4228afa957504d92899a16b18f1db/5C599844/t51.2885-15/sh0.08/e35/s750x750/41013725_540667323027121_5550989628702074250_n.jpg", 750, 750},
				{"https://instagram.fist1-2.fna.fbcdn.net/vp/bfcf8f6c08868abc8f6afc1d9d6d43b0/5C5225A1/t51.2885-15/e35/41013725_540667323027121_5550989628702074250_n.jpg", 1080, 1080},
			},
		},
	}
	edges, _ := p.sidecarEdges()
	if len(edges) != len(cases) {
		t.Errorf("number of edge nodes: got: %d, want: %d", len(edges), len(cases))
	}
	for i, c := range cases {
		got, _ := edges[i].displayResources()
		if !reflect.DeepEqual(c.res, got) {
			t.Errorf("got: %q, want: %q", got, c.res)
		}
	}
}

func TestProfilePage(t *testing.T) {
	p := new(ProfilePostPage)
	files := []string{
		"json/instagram-profile-page.json",
		"json/instagram-single-pic.json",
		"json/instagram-multi-pic.json",
	}
	for _, f := range files {
		buf := readAllOrDie(f, t)
		err := json.Unmarshal(buf, p)
		if err != nil {
			t.Error(err)
		}
	}
}

type testTimeSource struct {
	t time.Time
}

func (t testTimeSource) time() time.Time {
	return t.t
}

var dates = []struct {
	tm   testTimeSource
	want rangeStatus
}{

	// time values in reverse chronological order
	{testTimeSource{time.Date(2006, time.Month(11), 3, 10, 28, 0, 0, time.UTC)}, cont},
	{testTimeSource{time.Date(2005, time.Month(12), 2, 7, 25, 0, 0, time.UTC)}, cont},
	{testTimeSource{time.Date(2005, time.Month(11), 21, 22, 15, 0, 0, time.UTC)}, inRange},
	{testTimeSource{time.Date(2005, time.Month(11), 3, 10, 28, 0, 0, time.UTC)}, inRange},
	{testTimeSource{time.Date(2004, time.Month(10), 3, 10, 27, 0, 0, time.UTC)}, outOfRange},
	{testTimeSource{time.Date(2004, time.Month(8), 7, 2, 23, 0, 0, time.UTC)}, outOfRange},
}

func TestTimeRange(t *testing.T) {
	r := timeRange{
		start: time.Date(2005, time.Month(11), 21, 22, 15, 0, 0, time.UTC),
		end:   time.Date(2004, time.Month(10), 3, 10, 27, 0, 0, time.UTC),
	}
	for _, d := range dates {
		got := r.includes(d.tm)
		if got != d.want {
			t.Errorf("got: %v, want: %v", got, d.want)
		}
	}
}

type ts struct{}

func (ts) time() time.Time {
	return time.Now()
}

func TestCounRange(t *testing.T) {
	expected := []rangeStatus{
		cont, cont, cont, inRange, inRange,
		inRange, inRange, outOfRange, outOfRange,
	}
	var ts ts
	cr := countRange{
		off:   3,
		count: 4,
	}
	for _, want := range expected {
		got := cr.includes(ts)
		if got != want {
			t.Errorf("got: %v, want: %v, s: %v", got, want, cr)
		}
	}
}

func TestCountTimeRange(t *testing.T) {
	cr := countTimeRange{
		off: 2,
		to:  time.Date(2004, time.Month(10), 3, 10, 27, 0, 0, time.UTC),
	}
	for _, d := range dates {
		if got := cr.includes(d.tm); got != d.want {
			t.Errorf("got: %v, want: %v", got, d.want)
		}
	}
}

func TestTimeCountRange(t *testing.T) {
	cr := timeCountRange{
		from:  time.Date(2005, time.Month(11), 21, 22, 15, 0, 0, time.UTC),
		count: 2,
	}
	for _, d := range dates {
		if got := cr.includes(d.tm); got != d.want {
			t.Errorf("got: %v, want: %v, tim: %v", got, d.want, d.tm.time())
		}
	}
}
