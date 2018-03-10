package main

import (
	"fmt"
	"net/http"
)

type ShortenHandler struct {
	urlMap map[string]string
}

func main() {
	http.ListenAndServe(":5050", &ShortenHandler{
		urlMap: readMapping(),
	})
}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI[1:]
	fmt.Println(path)

	url, ok := s.urlMap[path]

	if ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	fmt.Fprintln(w, "Path '"+path+"' not mapped !")
}

func readMapping() map[string]string {
	return map[string]string{
		"test": "http://localhost:5050/popo",
		"popo": "http://localhost:5050/toto",
		"toto": "http://localhost:5050/end",
	}
}
