package main

import (
	"comic-rss-backend/handlers"
	"comic-rss-backend/models"
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
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
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

	// クエリの定義
	articleID := 1
	const sqlStr = `
		select *
		from articles
		where article_id = $1;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime

		err := rows.Scan(&article.ID, &article.Title, &article.Contents,
			&article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		articleArray = append(articleArray, article)
	}

	fmt.Printf("%+v\n", articleArray)
}
