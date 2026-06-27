package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddleware {
		n = globalMiddleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}
