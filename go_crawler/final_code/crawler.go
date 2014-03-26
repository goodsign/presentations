package main

import (
	"fmt"
	"github.com/PuerkitoBio/gocrawl"
	"math/rand"
	"time"
)

// Create the Extender implementation, based on the gocrawl-provided DefaultExtender,
// because we don't want/need to override all methods.
type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visited() and Filter()
}

func randomColorHex() string {
	rn := func() int {
		return 80 + rand.Int()%160
	}
	return fmt.Sprintf("#%.2x%.2x%.2x", rn(), rn(), rn())
}

// Override Filter for our need.
func (this *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}

func (ext *ExampleExtender) Visited(ctx *gocrawl.URLContext, harvested interface{}) {
	newStr := ctx.NormalizedURL().Path
	_, ok := CrawlerSiteGraph.Nodes[newStr]
	if !ok {
		CrawlerSiteGraph.Nodes[newStr] = SiteGraphNode{
			Label: newStr,
			Color: randomColorHex(),
			Mass:  1.0,
		}
	}

	sourceUrl := ctx.NormalizedSourceURL()
	if sourceUrl != nil {
		sourceStr := sourceUrl.Path
		if _, ok := CrawlerSiteGraph.Edges[sourceStr]; !ok {
			CrawlerSiteGraph.Edges[sourceStr] = SiteGraphEdge{}
		}
		CrawlerSiteGraph.Edges[sourceStr][newStr] = SiteGraphEdgeRelation{
			Directed: true,
		}
	}

	CrawlerSiteGraph.NodeCount = len(CrawlerSiteGraph.Nodes)
}

func initCrawler(maxVisits int) *gocrawl.Crawler {
	opts := gocrawl.NewOptions(&ExampleExtender{gocrawl.DefaultExtender{}})
	opts.CrawlDelay = 500 * time.Millisecond
	opts.LogFlags = gocrawl.LogInfo
	opts.MaxVisits = maxVisits
	return gocrawl.NewCrawlerWithOptions(opts)
}

func StartCrawler(url string, maxVisits int) {
	cr := initCrawler(maxVisits)
	stopChan := make(chan int)
	go PulseBeat(stopChan)
	cr.Run(url)
	stopChan <- 0
}
