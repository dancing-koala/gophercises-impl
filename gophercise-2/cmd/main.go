package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type urlMap map[string]string

type ShortenHandler struct {
	mapping urlMap
}

func main() {
	var src string
	var jsonSrc bool

	flag.StringVar(&src, "map-src", "./mapping.json", "Path of the file containing the mappings")
	flag.BoolVar(&jsonSrc, "json", false, "Indicates whether the source file is a json file")

	closeChan := make(chan interface{}, 1)

	go func(um urlMap) {
		http.ListenAndServe(":5050", &ShortenHandler{
			mapping: um,
		})

		closeChan <- nil
	}(readMapping(src, jsonSrc))

	<-closeChan
}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI[1:]
	url, ok := s.mapping[path]

	if ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Path %q not mapped !\n", path)
}

func readMapping(src string, isJson bool) urlMap {
	fmt.Println("src="+src, "isJson="+strconv.FormatBool(isJson))

	data, err := ioutil.ReadFile(src)

	if err != nil {
		panic(err)
	}

	var parsedResult map[string]string

	err = json.Unmarshal(data, &parsedResult)

	if err != nil {
		panic(err)
	}

	return parsedResult
}
