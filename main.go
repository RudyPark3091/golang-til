package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("default: ", r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("param: ", r.Form["test_param"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Golang Webserver Working!")
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "public/index.html"
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	default:
		contentType = "text/plain"
	}

	return contentType
}

func main() {
	PORT := 8080

	http.HandleFunc("/", defaultHandler)
	http.Handle("/html", new(staticHandler))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Printf("Server Started -> Port %d", PORT)
	}
}
