package main

import (
	"github.com/PuerkitoBio/gocrawl"
	"time"
)

// Create the Extender implementation, based on the gocrawl-provided DefaultExtender,
// because we don't want/need to override all methods.
type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visited() and Filter()
}

// Override Filter for our need.
func (this *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}

func (ext *ExampleExtender) Visited(ctx *gocrawl.URLContext, harvested interface{}) {

}

func initCrawler(maxVisits int) *gocrawl.Crawler {
	opts := gocrawl.NewOptions(&ExampleExtender{gocrawl.DefaultExtender{}})
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll
	opts.MaxVisits = maxVisits
	return gocrawl.NewCrawlerWithOptions(opts)
}
