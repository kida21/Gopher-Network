package main

import (
	"log"
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter , r *http.Request,err error) {
    log.Printf("internal error :%s on path:%s error:%s",r.Method,r.URL.Path,err)
	WriteError(w,http.StatusInternalServerError,"the serever encountered a problem")
}
func (app *application) BadRequestResponse(w http.ResponseWriter,r *http.Request,err error) {
	 log.Printf("Bad request error :%s on path:%s error:%s",r.Method,r.URL.Path,err)
	WriteError(w,http.StatusBadRequest,err.Error())
}
func (app *application) NotFoundResponse(w http.ResponseWriter,r *http.Request,err error) {
	 log.Printf("Not Found error :%s on path:%s error:%s",r.Method,r.URL.Path,err)
	WriteError(w,http.StatusNotFound,"Not Found")
}