package main

import (
	"html/template"
	"log"
	"net/http"
)

const ADDR = ":9090"

var l *log.Logger
var templates map[string]*template.Template

func init() {
	// init
	l = log.Default()
	templates = make(map[string]*template.Template, 1)

	//load template
	templates["main"] = template.Must(template.ParseFiles("pages/page.html"))
}

func handlePages(stories Story) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		requestedPages, present := query["page"]
		if !present || len(requestedPages) == 0 {
			requestedPages = make([]string, 1, 1)
			requestedPages[0] = "intro"
		}
		// generate HTML page
		writer.Header().Set("Content-Type", "text/html")
		storyData := stories[requestedPages[0]]
		if err := templates["main"].Execute(writer, storyData); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	stories := Reader{
		filename: "gopher.json",
		logger:   l,
	}.Read()

	http.HandleFunc("/", handlePages(stories))
	log.Fatal(http.ListenAndServe(ADDR, nil))
}
