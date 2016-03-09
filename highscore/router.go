package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HttpRouter interface {
	GET(url string, handler func(w http.ResponseWriter, r *http.Request))
	Router() http.Handler
}

type GorillaRouter struct {
	router *mux.Router
}

func (r *GorillaRouter) GET(url string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.router.Handle(url, http.HandlerFunc(handler))
}

func NewGorillaRouter(router *mux.Router) HttpRouter {
	return &GorillaRouter{router}
}

func (r *GorillaRouter) Router() http.Handler {
	return r.router
}
