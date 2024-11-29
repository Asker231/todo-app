package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Asker231/todo-app.git/configs"
	"github.com/Asker231/todo-app.git/pkg/res"
	"github.com/go-playground/validator/v10"
)
type Auth struct{
	*configs.ConfigApp
}
type DepAuth struct{
	*configs.ConfigApp
}
func(a *Auth)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		//read body ...
			var payload LoginRequest
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil{
				res.JsonRes(w,err.Error(),402)
			}

			validate := validator.New()
			err = validate.Struct(payload)
			if err != nil{
				res.JsonRes(w,err.Error(),402)
				return
			}

				//write log ...
				result := LoginResponse{
					Token: a.ConfigApp.ConfAuth.TOKEN,
				}
				res.JsonRes(w,result,200)
	}
}

func(a *Auth)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func AuthHandler(router *http.ServeMux,deps DepAuth){
	auth := &Auth{
		ConfigApp: deps.ConfigApp,
	}
	router.HandleFunc("POST /auth/login",auth.Login())
	router.HandleFunc("POST /auth/register",auth.Register())
}
