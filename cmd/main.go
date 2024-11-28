package main

import (
	"net/http"
	"github.com/Asker231/todo-app.git/internal/auth"
)


func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	auth.AuthHandler(router)
	server.ListenAndServe()
}