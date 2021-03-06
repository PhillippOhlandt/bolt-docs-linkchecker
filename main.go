package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

var (
	host                  = "https://docs.bolt.cm"
	versions              = []string{"2.2", "3.0"}
	currentCrawledVersion = "2.2"
)

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//fmt.Printf("[%v] %s\n", res.StatusCode, ctx.URL())
	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}

	return true
}

func (e *Ext) RequestGet(ctx *gocrawl.URLContext, headRes *http.Response) bool {
	if headRes.StatusCode >= 200 && headRes.StatusCode < 300 {
		return true
	}

	versionUrlPart := fmt.Sprintf("/%s/", currentCrawledVersion)
	sourceUrl := fmt.Sprint(ctx.SourceURL())
	// When request returns 4xx or 5xx response and the source URL contains our current crawled version, mark as error.
	if strings.Contains(sourceUrl, versionUrlPart) {
		fmt.Printf("\x1b[31;1m[%v]\x1b[0m %s from \x1b[31;1m%s\x1b[0m\n", headRes.StatusCode, ctx.URL(), ctx.SourceURL())
	}

	return false
}

func main() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Millisecond
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = true
	opts.MaxVisits = 1000000
	opts.HeadBeforeGet = true

	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	if len(os.Args) > 2 {
		versions = strings.Split(os.Args[2], ",")
	}

	c := gocrawl.NewCrawlerWithOptions(opts)

	for _, version := range versions {
		url := fmt.Sprintf("%s/%s/", host, version)
		currentCrawledVersion = version
		fmt.Println("Crawling ", url)
		c.Run(url)
	}
}
