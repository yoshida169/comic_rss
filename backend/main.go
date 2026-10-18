package main

import (
	"comic-rss-backend/db"
	"comic-rss-backend/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// dbDSN returns the MySQL DSN. DB_DSN takes precedence; otherwise it is built
// from USERNAME / USERPASS / DATABASE in the repository-root .env file
// (the same variables docker-compose.yaml uses). DB_HOST overrides the
// default host:port of 127.0.0.1:3306.
func dbDSN() (string, error) {
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		return dsn, nil
	}

	// Read into a map instead of Load so that shell variables such as zsh's
	// USERNAME cannot shadow the values in .env.
	env, err := godotenv.Read("../.env")
	if err != nil {
		return "", fmt.Errorf("DB_DSN is not set and .env could not be read: %w", err)
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1:3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		env["USERNAME"], env["USERPASS"], host, env["DATABASE"]), nil
}

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
	dsn, err := dbDSN()
	if err != nil {
		log.Fatal(err)
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
