package main

import (
	"HttpServer/configs"
	"HttpServer/internal/auth"
	"HttpServer/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	auth.NewAuth(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
