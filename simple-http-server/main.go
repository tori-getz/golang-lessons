package main

import (
	"fmt"
	"net/http"
	"simple-http-server/app/hello"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server listening at 3000 port")
	server.ListenAndServe()
}
