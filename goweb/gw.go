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
	routerGroups []*routerGroup
}

type routerGroup struct {
	name           string
	handlerFuncMap map[string]HandlerFunc
}

func (r *router) Group(name string) *routerGroup {
	rg := &routerGroup{
		name:           name,
		handlerFuncMap: make(map[string]HandlerFunc),
	}
	r.routerGroups = append(r.routerGroups, rg)
	return rg
}

func (rg *routerGroup) Add(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
}

func New() *Engine {
	return &Engine{
		router: router{},
	}
}

func (e *Engine) Run() {
	// user key:get value:func
	for _, group := range e.routerGroups {
		for addr, handlerFunc := range group.handlerFuncMap {
			fmt.Println("/" + group.name + addr)
			http.HandleFunc("/"+group.name+addr, handlerFunc)
		}
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
