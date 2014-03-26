package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fUrl       = flag.String("u", "", "Initial URL for crawler")
	fMaxVisits = flag.Int("maxv", 100, "Max page visits for crawler")
	fPort      = flag.Int("p", 8080, "Web UI port")
)

func validateInput() error {
	if len(*fUrl) == 0 {
		return fmt.Errorf("initial URL cannot be nil")
	}
	return nil
}

func main() {
	flag.Parse()
	if err := validateInput(); err != nil {
		fmt.Printf("Invalid input params: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Starting crawler. Initial URL: '%s', max visits: %d\n", *fUrl, *fMaxVisits)
	go StartCrawler(*fUrl, *fMaxVisits)
	ServeHttp(*fPort)
}
