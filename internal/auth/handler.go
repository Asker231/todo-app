package auth

import "net/http"

type Auth struct{}

func(a *Auth)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func(a *Auth)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func AuthHandler(router *http.ServeMux){
	auth := Auth{}
	router.HandleFunc("/auth/login",auth.Login())
	router.HandleFunc("/auth/register",auth.Register())
}
