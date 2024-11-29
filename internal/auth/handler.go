package auth

import (
	"fmt"
	"github.com/Asker231/todo-app.git/configs"
	"github.com/Asker231/todo-app.git/pkg/req"
	"github.com/Asker231/todo-app.git/pkg/res"
	"net/http"
)

type Auth struct {
	*configs.ConfigApp
}
type DepAuth struct {
	*configs.ConfigApp
}

func (a *Auth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		token := LoginResponse{
			Token: a.ConfigApp.ConfAuth.TOKEN,
		}
		res.JsonRes(w, token, 200)
	}
}

func (a *Auth) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		token := RegisterResponse{
			Token: a.ConfigApp.ConfAuth.TOKEN,
		}
		res.JsonRes(w, token, 200)
	}
}

func AuthHandler(router *http.ServeMux, deps DepAuth) {
	auth := &Auth{
		ConfigApp: deps.ConfigApp,
	}
	router.HandleFunc("POST /auth/login", auth.Login())
	router.HandleFunc("POST /auth/register", auth.Register())
}
