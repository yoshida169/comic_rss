package handlers

import (
	"comic-rss-backend/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); got != "Hello world\n" {
		t.Fatalf("body = %q, want %q", got, "Hello world\n")
	}
}

func TestDBHealthHandler(t *testing.T) {
	t.Run("DB未設定の場合は503", func(t *testing.T) {
		old := DB
		DB = nil
		defer func() { DB = old }()

		req := httptest.NewRequest(http.MethodGet, "/health/db", nil)
		rec := httptest.NewRecorder()

		DBHealthHandler(rec, req)

		if rec.Code != http.StatusServiceUnavailable {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
		}
	})
}

func TestPostArticleHandler(t *testing.T) {
	t.Run("有効なJSONはそのまま返す", func(t *testing.T) {
		body := `{"article_id":10,"title":"new","contents":"body","user_name":"saki"}`
		req := httptest.NewRequest(http.MethodPost, "/article", strings.NewReader(body))
		rec := httptest.NewRecorder()

		PostArticleHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var got models.Article
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if got.ID != 10 || got.Title != "new" || got.Contents != "body" || got.UserName != "saki" {
			t.Fatalf("unexpected article: %+v", got)
		}
	})

	t.Run("ボディなしは400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article", nil)
		rec := httptest.NewRecorder()

		PostArticleHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestArticleDetailHandler(t *testing.T) {
	t.Run("記事詳細をJSONで返す", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rec := httptest.NewRecorder()

		ArticleDetailHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var got models.Article
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if got.ID != models.Article1.ID || got.Title != models.Article1.Title {
			t.Fatalf("unexpected article: %+v", got)
		}
	})

	t.Run("idが数値でない場合は400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})
		rec := httptest.NewRecorder()

		ArticleDetailHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestArticleListHandler(t *testing.T) {
	t.Run("記事一覧をJSONで返す", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/list", nil)
		rec := httptest.NewRecorder()

		ArticleListHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var got []models.Article
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if len(got) != 2 || got[0].ID != 1 || got[1].ID != 2 {
			t.Fatalf("unexpected list: %+v", got)
		}
	})

	t.Run("pageが数値でない場合は400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/list?page=abc", nil)
		rec := httptest.NewRecorder()

		ArticleListHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestPostNiceHandler(t *testing.T) {
	t.Run("有効なJSONはそのまま返す", func(t *testing.T) {
		body := `{"article_id":1,"nice":5}`
		req := httptest.NewRequest(http.MethodPost, "/article/nice", strings.NewReader(body))
		rec := httptest.NewRecorder()

		PostNiceHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var got models.Article
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if got.ID != 1 || got.NiceNum != 5 {
			t.Fatalf("unexpected article: %+v", got)
		}
	})

	t.Run("ボディなしは400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article/nice", nil)
		rec := httptest.NewRecorder()

		PostNiceHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestPostCommentHandler(t *testing.T) {
	t.Run("有効なJSONはそのまま返す", func(t *testing.T) {
		body := `{"comment_id":3,"article_id":1,"message":"hi"}`
		req := httptest.NewRequest(http.MethodPost, "/comment", strings.NewReader(body))
		rec := httptest.NewRecorder()

		PostCommentHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var got models.Comment
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if got.CommentID != 3 || got.ArticleID != 1 || got.Message != "hi" {
			t.Fatalf("unexpected comment: %+v", got)
		}
	})

	t.Run("ボディなしは400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/comment", nil)
		rec := httptest.NewRecorder()

		PostCommentHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}
