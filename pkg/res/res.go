package res

import (
	"encoding/json"
	"net/http"
)
func JsonRes(resp http.ResponseWriter,data any,status int){
	resp.Header().Set("Content-Type","application/json")
	resp.WriteHeader(status)
	json.NewEncoder(resp).Encode(data)
}