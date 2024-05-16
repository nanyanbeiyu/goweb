package main

import (
	"fmt"
	"github.com/nanyanbeiyu/goweb"
	"net/http"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "Hello word")
	//})
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	gw := goweb.New()
	user := gw.Group("user")
	user.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello word")
	})
	gw.Run()
}
