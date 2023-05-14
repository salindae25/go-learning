package main

import (
	"fmt"
	"log"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

// Router is array of Route struct
type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(path string, method string, handler http.Handler) {
	r.routes = append(r.routes, Route{Path: path, Method: method, Handler: handler})
}
func (r *Router) Get(path string, handler http.Handler) {
	r.AddRoute(path, "GET", handler)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.AddRoute(path, "POST", handler)
}

func (r *Router) Put(path string, handler http.Handler) {
	r.AddRoute(path, "PUT", handler)
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.AddRoute(path, "DELETE", handler)
}
func (r *Router) getHandler(path string, method string) http.Handler {
	for _, v := range r.routes {
		if v.Path == path && v.Method == method {
			return v.Handler
		}
	}
	return http.NotFoundHandler()
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	handler := r.getHandler(path, method)

	handler.ServeHTTP(w, req)
}
func helloHadler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
}
func main() {
	mux := NewRouter()
	mux.Get("/hello", helloHadler())
	mux.Get("/hello/:id", helloHadler())
	mux.Post("/hello", helloHadler())
	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Panic(err)
	}
}
