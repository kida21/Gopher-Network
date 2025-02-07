package main

import (

	"net/http"
)

func (app *application) HealthCheckHanlder(w http.ResponseWriter,r *http.Request){
	data := map[string]string{
		"status":"ok",
		"env":app.config.env,
		"version":version,
	}
	if err :=WriteJson(w,http.StatusOK,data);err !=nil{
		app.InternalServerError(w,r,err)
	}
}