package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/PuerkitoBio/gocrawl"
    "github.com/PuerkitoBio/goquery"
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
    if ctx.URL().Host == "docs.bolt.cm"{
        return true
    }
    return false
}

func (e *Ext) RequestGet(ctx *gocrawl.URLContext, headRes *http.Response) bool {
    if headRes.StatusCode >= 200 && headRes.StatusCode < 300 {
        return true
    }
    
    fmt.Printf("[%v] %s from %s\n", headRes.StatusCode, ctx.URL(), ctx.SourceURL())
    return false
}

func main() {
    ext := &Ext{&gocrawl.DefaultExtender{}}
    // Set custom options
    opts := gocrawl.NewOptions(ext)
    opts.CrawlDelay = 2 * time.Millisecond
    opts.LogFlags = gocrawl.LogError
    opts.SameHostOnly = false
    opts.MaxVisits = 1000000
    opts.HeadBeforeGet = true

    c := gocrawl.NewCrawlerWithOptions(opts)
    c.Run("https://docs.bolt.cm/2.2/")
    c.Run("https://docs.bolt.cm/3.0/")
    c.Run("https://docs.bolt.cm/3.1/")
}