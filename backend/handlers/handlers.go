package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world\n")
}

func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

func(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}