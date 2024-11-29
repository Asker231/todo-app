package req

import (
	"github.com/Asker231/todo-app.git/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	data, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}
	err = IsValid(data)
	if err != nil {
		res.JsonRes(w, err.Error(), 402)
		return nil, err
	}
	return data, nil
}
