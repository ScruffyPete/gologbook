package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareMux(t *testing.T) {
	t.Run("should apply middlewares to the request", func(t *testing.T) {
		middleware := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
				w.Write([]byte("I'm a teapot"))
				h.ServeHTTP(w, r)
			})
		}
		mux := NewMiddlewareMux(middleware)
		mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusTeapot, w.Code)
		assert.Equal(t, "I'm a teapot", w.Body.String())
	})

	t.Run("should apply middlewares in reverse order", func(t *testing.T) {
		firstMiddleware := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
				w.Write([]byte("First middleware\n"))
				h.ServeHTTP(w, r)
			})
		}

		secondMiddleware := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Second middleware\n"))
				h.ServeHTTP(w, r)
			})
		}

		mux := NewMiddlewareMux(secondMiddleware, firstMiddleware)
		mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Second middleware\nFirst middleware\n", w.Body.String())
	})
}
