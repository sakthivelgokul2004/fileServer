package handlers

import (
	"log"
	"net/http"
)

func ServecFiles(w http.ResponseWriter, r *http.Request) {
	path := "./static" + r.URL.Path
	if path == "./static/" {
		path += "index.html"
	}
	log.Printf("path is %v",path)
	http.ServeFile(w, r, path)
}
