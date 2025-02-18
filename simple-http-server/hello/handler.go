package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}

	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello request", r)
		w.Write([]byte("Hello"))
	}
}
