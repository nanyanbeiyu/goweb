package goweb

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	handlerFuncMap map[string]HandlerFunc
}

func (r *router) Add(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
}

func New() *Engine {
	return &Engine{
		router: router{
			handlerFuncMap: make(map[string]HandlerFunc),
		},
	}
}

func (e *Engine) Run() {
	for addr, handlerFunc := range e.router.handlerFuncMap {
		http.HandleFunc(addr, handlerFunc)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
