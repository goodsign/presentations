package main

import "encoding/json"

type SiteGraphNode struct {
	Label string  `json:"label"`
	Color string  `json:"color"`
	Mass  float32 `json:"mass"`
}

type SiteGraphEdgeRelation struct {
	Directed bool `json:"directed"`
}

type SiteGraphEdge map[string]SiteGraphEdgeRelation

type SiteGraph struct {
	Finished  bool                     `json:"finished"`
	NodeCount int                      `json:"nodeCount"`
	Nodes     map[string]SiteGraphNode `json:"nodes"`
	Edges     map[string]SiteGraphEdge `json:"edges"`
}

func (g SiteGraph) String() string {
	b, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(b)
}

var crawlerNodes = map[string]SiteGraphNode{}
var crawlerEdges = map[string]SiteGraphEdge{}

var CrawlerSiteGraph = &SiteGraph{
	Finished:  false,
	NodeCount: 0,
	Nodes:     crawlerNodes,
	Edges:     crawlerEdges,
}
