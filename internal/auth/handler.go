package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/todo-app.git/configs"
)

type Auth struct{
	*configs.ConfigApp
}

type DepAuth struct{
	*configs.ConfigApp
}

func(a *Auth)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Yor token is:%s\n",a.ConfigApp.TOKEN)
	}
}

func(a *Auth)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func AuthHandler(router *http.ServeMux,deps DepAuth){
	auth := Auth{
		ConfigApp: deps.ConfigApp,
	}
	router.HandleFunc("/auth/login",auth.Login())
	router.HandleFunc("/auth/register",auth.Register())
}
