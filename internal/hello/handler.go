package hello

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
}

func NewHelloHandler(router *http.ServeMux) {
	handler := &helloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *helloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hello ooooo")
	}
}
