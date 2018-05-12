package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ArcOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type ArcStory struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []ArcOption `json:"options"`
}

type ArcStories map[string]ArcStory

type ArcHandler struct{}

var (
	stories = make(ArcStories)
	tpl     *template.Template
)

func main() {
	rawTemplate, err := ioutil.ReadFile("./story.tpl")

	handleErr(err)

	tpl, err = template.New("webpage").Parse(string(rawTemplate))

	handleErr(err)

	data, err := ioutil.ReadFile("./story.json")

	handleErr(err)

	err = json.Unmarshal(data, &stories)

	handleErr(err)

	http.ListenAndServe(":9999", &ArcHandler{})
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (ah *ArcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.RequestURI[1:]

	item, ok := stories[slug]

	fmt.Println(slug, item, ok)

	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "text/html")

	err := tpl.Execute(w, item)

	handleErr(err)
}
