package main

import (
	"comic-rss-backend/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/{id: [0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	return r
}

func main() {
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
