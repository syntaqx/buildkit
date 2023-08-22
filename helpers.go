package buildkit

import "net/http"

func (e *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodGet, pattern, handlerFunc)
}

func (e *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodPost, pattern, handlerFunc)
}

func (e *Engine) Put(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodPut, pattern, handlerFunc)
}

func (e *Engine) Delete(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodDelete, pattern, handlerFunc)
}

func (e *Engine) Patch(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodPatch, pattern, handlerFunc)
}

func (e *Engine) Head(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodHead, pattern, handlerFunc)
}

func (e *Engine) Options(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute(http.MethodOptions, pattern, handlerFunc)
}
