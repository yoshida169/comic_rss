package main

import (
	"comic-rss-backend/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

// type Comment struct {
// 	CommentID int       `json:"comment_id"`
// 	ArticleID int       `json:"article_id"`
// 	Message   string    `json:"message"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type Article struct {
// 	ID          int       `json:"article_id"`
// 	Title       string    `json:"title"`
// 	Contents    string    `json:"contents"`
// 	UserName    string    `json:"user_name"`
// 	NiceNum     int       `json:"nice"`
// 	CommentList []Comment `json:"comments"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

func main() {

	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
