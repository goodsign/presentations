package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	TemplatesFolderPath = "templates"
	HomeTemplate        = "home.html"
)

func renderTemplate(name string, data interface{}, w http.ResponseWriter) error {
	t, err := template.ParseFiles(filepath.Join(TemplatesFolderPath, name))
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(HomeTemplate, nil, w)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err)
	}
}

func GraphJsonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, CrawlerSiteGraph)
}

func ServeHttp(port int) {
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/graph.json", GraphJsonHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
