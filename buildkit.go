package buildkit

import (
	"context"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type MiddlewareFunc func(http.Handler) http.Handler

type route struct {
	pattern     string
	method      string
	handlerFunc HandlerFunc
}

type Engine struct {
	routes      []route
	middlewares []MiddlewareFunc
	router      *trie
}

func New() *Engine {
	return &Engine{
		router: &trie{root: newNode()},
	}
}

func (e *Engine) Use(middleware MiddlewareFunc) {
	e.middlewares = append(e.middlewares, middleware)
}

func (e *Engine) AddRoute(method, pattern string, handlerFunc HandlerFunc) {
	r := route{pattern: pattern, method: method, handlerFunc: handlerFunc}
	e.routes = append(e.routes, r)
	e.router.insert(pattern, &r)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var finalHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeMatch := e.router.search(r.URL.Path)
		if routeMatch != nil && r.Method == routeMatch.route.method {
			if len(routeMatch.params) > 0 {
				ctx := context.WithValue(r.Context(), paramKey, routeMatch.params)
				r = r.WithContext(ctx)
			}
			routeMatch.route.handlerFunc(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// Apply middlewares in order of addition
	for _, middleware := range e.middlewares {
		finalHandler = middleware(finalHandler)
	}

	finalHandler.ServeHTTP(w, r)
}

// GetParams retrieves all dynamic route parameters from the request context.
func GetParams(r *http.Request) map[string]string {
	val := r.Context().Value(paramKey)
	if val == nil {
		return nil
	}
	return val.(map[string]string)
}

// GetParam retrieves the value of a dynamic route parameter by its name.
// If the parameter does not exist, it returns an empty string.
func GetParam(r *http.Request, name string) string {
	params := GetParams(r)
	if params == nil {
		return ""
	}
	return params[name]
}
