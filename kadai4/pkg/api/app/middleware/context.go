package middleware

import (
	"context"
	"net/http"
	"time"
)

type MiddleWare interface {
	ServeNext(h http.Handler) http.Handler
}

type MiddleWareFunc func(h http.Handler) http.Handler

func (f MiddleWareFunc) ServeNext(h http.Handler) http.Handler {
	return f(h)
}

func With(h http.Handler, ms ...MiddleWare) http.Handler {
	for _, m := range ms {
		h = m.ServeNext(h)
	}
	return h
}

type ContextMiddleWare struct{}

func (m ContextMiddleWare) ServeNext(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), "time", time.Now())
		r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

type ResponseHeaderMiddleWare struct{}

func (m ResponseHeaderMiddleWare) ServeNext(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
