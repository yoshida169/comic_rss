package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		body       string
		wantStatus int
		wantBody   string // 空でなければ完全一致で比較
	}{
		{
			name:       "GET /hello",
			method:     http.MethodGet,
			path:       "/hello",
			wantStatus: http.StatusOK,
			wantBody:   "Hello world\n",
		},
		{
			name:       "POST /hello はメソッド不一致で405",
			method:     http.MethodPost,
			path:       "/hello",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "GET /health/db はDB未設定で503",
			method:     http.MethodGet,
			path:       "/health/db",
			wantStatus: http.StatusServiceUnavailable,
		},
		{
			name:       "POST /article (JSONあり)",
			method:     http.MethodPost,
			path:       "/article",
			body:       `{"title":"t"}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "POST /article (ボディなし)",
			method:     http.MethodPost,
			path:       "/article",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "GET /article/list",
			method:     http.MethodGet,
			path:       "/article/list",
			wantStatus: http.StatusOK,
		},
		{
			name:       "GET /article/list?page=abc",
			method:     http.MethodGet,
			path:       "/article/list?page=abc",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "GET /article/1",
			method:     http.MethodGet,
			path:       "/article/1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "POST /article/1 はメソッド不一致で405",
			method:     http.MethodPost,
			path:       "/article/1",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "POST /article/nice (JSONあり)",
			method:     http.MethodPost,
			path:       "/article/nice",
			body:       `{"article_id":1}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "POST /comment (JSONあり)",
			method:     http.MethodPost,
			path:       "/comment",
			body:       `{"message":"hi"}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "POST /login (不正なJSON)",
			method:     http.MethodPost,
			path:       "/login",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "GET /login はメソッド不一致で405",
			method:     http.MethodGet,
			path:       "/login",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	router := newRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			} else {
				req = httptest.NewRequest(tt.method, tt.path, nil)
			}
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantBody != "" {
				if got := rec.Body.String(); got != tt.wantBody {
					t.Fatalf("body = %q, want %q", got, tt.wantBody)
				}
			}
		})
	}
}

func TestRouterNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}
