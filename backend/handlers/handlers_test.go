package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("GET returns hello message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)
		rec := httptest.NewRecorder()

		HelloHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Hello world\n" {
			t.Fatalf("body = %q, want %q", got, "Hello world\n")
		}
	})

	t.Run("POST returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/hello", nil)
		rec := httptest.NewRecorder()

		HelloHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}

func TestPostArticleHandler(t *testing.T) {
	t.Run("POST returns posting message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article", nil)
		rec := httptest.NewRecorder()

		PostArticleHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Posting Article...\n" {
			t.Fatalf("body = %q, want %q", got, "Posting Article...\n")
		}
	})

	t.Run("GET returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article", nil)
		rec := httptest.NewRecorder()

		PostArticleHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}

func TestArticleDetailHandler(t *testing.T) {
	t.Run("GET returns article detail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/1", nil)
		rec := httptest.NewRecorder()

		ArticleDetailHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Article No.1\n" {
			t.Fatalf("body = %q, want %q", got, "Article No.1\n")
		}
	})

	t.Run("POST returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article/1", nil)
		rec := httptest.NewRecorder()

		ArticleDetailHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}

func TestArticleListHandler(t *testing.T) {
	t.Run("GET returns article list", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/list", nil)
		rec := httptest.NewRecorder()

		ArticleListHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Article List\n" {
			t.Fatalf("body = %q, want %q", got, "Article List\n")
		}
	})

	t.Run("POST returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article/list", nil)
		rec := httptest.NewRecorder()

		ArticleListHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}

func TestPostNiceHandler(t *testing.T) {
	t.Run("POST returns posting nice message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/article/nice", nil)
		rec := httptest.NewRecorder()

		PostNiceHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Posting Nice...\n" {
			t.Fatalf("body = %q, want %q", got, "Posting Nice...\n")
		}
	})

	t.Run("GET returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/nice", nil)
		rec := httptest.NewRecorder()

		PostNiceHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}

func TestPostCommentHandler(t *testing.T) {
	t.Run("POST returns posting comment message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/comment", nil)
		rec := httptest.NewRecorder()

		PostCommentHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Body.String(); got != "Posting Comment...\n" {
			t.Fatalf("body = %q, want %q", got, "Posting Comment...\n")
		}
	})

	t.Run("GET returns method not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comment", nil)
		rec := httptest.NewRecorder()

		PostCommentHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
		}
	})
}
