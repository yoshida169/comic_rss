package main

import (
	"comic-rss-backend/db"
	"comic-rss-backend/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func newRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodGet)
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
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set (e.g. user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true)")
	}

	database, err := db.Open(dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	handlers.SetDB(database)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
