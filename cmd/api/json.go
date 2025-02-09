package main

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
	
)

var Validate *validator.Validate
func init(){
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJson(w http.ResponseWriter,status int, data any)error {
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(status)
  return json.NewEncoder(w).Encode(data)
}

func ReadJson(w http.ResponseWriter,r *http.Request,data any) error{
	maxBytes := 1048578
    r.Body = http.MaxBytesReader(w,r.Body,int64(maxBytes))
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter ,status int,message string)error{
	type envelop struct{
         Error string `json:"error"`
	     }
	return WriteJson(w,status,&envelop{
		Error: message,
	})
	}
