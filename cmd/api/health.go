package main

import "net/http"

func (app *application) HealthCheckHanlder(w http.ResponseWriter,r *http.Request){
	w.Write([]byte ("Hello"))
}