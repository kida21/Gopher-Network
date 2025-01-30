package main

import (
	"net/http"

	"github.com/kida21/gopher/internal/store"
)

type CreatePayLoad struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}
func (app *application) createPostHandler(w http.ResponseWriter,r *http.Request) {
   var payload CreatePayLoad
   if err := ReadJson(w,r,&payload);err != nil{
	WriteError(w,http.StatusBadRequest,string(err.Error()))
   }
   ctx := r.Context()
   post := &store.Post{
	Title: payload.Title,
	Content: payload.Content,
	Tags: payload.Tags,
	UserID: 1,
   }
    if err := app.store.Posts.Create(ctx,post);err != nil{
		WriteError(w,http.StatusInternalServerError,err.Error())
		return
	}
	if err:=WriteJson(w,http.StatusCreated,post);err !=nil{
		WriteError(w,http.StatusInternalServerError,string(err.Error()))
		return
	}
}
