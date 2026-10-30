package main

import (
	"comic-rss-backend/handlers"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/health/db", handlers.DBHealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/{id: [0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)
	r.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)

	return r
}

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}
