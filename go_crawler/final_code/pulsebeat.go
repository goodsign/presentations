package main

import (
	"fmt"
	"time"
)

func PulseBeat(stop chan int) {
	ticker := time.Tick(time.Second)
	for {
		select {
		case <-stop:
			CrawlerSiteGraph.Finished = true
			return
		case t := <-ticker:
			fmt.Printf("-----------------\nTick: %s\n-----------------\n", t.Format("15:04:05"))
		}
	}
}
