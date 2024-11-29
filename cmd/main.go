package main

import (
	"net/http"

	"github.com/Asker231/todo-app.git/configs"
	"github.com/Asker231/todo-app.git/internal/auth"
	"github.com/Asker231/todo-app.git/pkg/db"
)

func main() {
	config := configs.CompileConfig()
	_ = db.ConnectionDb(config)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	auth.AuthHandler(router, auth.DepAuth{ConfigApp: config})
	server.ListenAndServe()
}
