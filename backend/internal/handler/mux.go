package handler

import "net/http"

type MiddlewareMux struct {
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func NewMiddlewareMux(middlewares ...func(http.Handler) http.Handler) *MiddlewareMux {
	mux := http.NewServeMux()
	return &MiddlewareMux{
		mux:         mux,
		middlewares: middlewares,
	}
}

func (mmux *MiddlewareMux) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	h := http.HandlerFunc(handler)
	wrappedHandler := mmux.applyMiddlewares(h)
	mmux.mux.Handle(path, wrappedHandler)
}

func (mmux *MiddlewareMux) applyMiddlewares(handler http.Handler) http.Handler {
	for i := len(mmux.middlewares) - 1; i >= 0; i-- {
		handler = mmux.middlewares[i](handler)
	}
	return handler
}

func (mmux *MiddlewareMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mmux.mux.ServeHTTP(w, r)
}
