package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "GET /hello",
			method:     http.MethodGet,
			path:       "/hello",
			wantStatus: http.StatusOK,
			wantBody:   "Hello world\n",
		},
		{
			name:       "POST /article",
			method:     http.MethodPost,
			path:       "/article",
			wantStatus: http.StatusOK,
			wantBody:   "Posting Article...\n",
		},
		{
			name:       "GET /article/list",
			method:     http.MethodGet,
			path:       "/article/list",
			wantStatus: http.StatusOK,
			wantBody:   "Article List\n",
		},
		{
			name:       "GET /article/1",
			method:     http.MethodGet,
			path:       "/article/1",
			wantStatus: http.StatusOK,
			wantBody:   "Article No.1\n",
		},
		{
			name:       "POST /article/nice",
			method:     http.MethodPost,
			path:       "/article/nice",
			wantStatus: http.StatusOK,
			wantBody:   "Posting Nice...\n",
		},
		{
			name:       "POST /comment",
			method:     http.MethodPost,
			path:       "/comment",
			wantStatus: http.StatusOK,
			wantBody:   "Posting Comment...\n",
		},
	}

	router := newRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if got := rec.Body.String(); got != tt.wantBody {
				t.Fatalf("body = %q, want %q", got, tt.wantBody)
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
