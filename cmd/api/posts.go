package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	app.BadRequestResponse(w,r,err)
   }
   ctx := r.Context()
   post := &store.Post{
	Title: payload.Title,
	Content: payload.Content,
	Tags: payload.Tags,
	UserID: 1,
   }
    if err := app.store.Posts.Create(ctx,post);err != nil{
       app.InternalServerError(w,r,err)
		return
	}
	if err:=WriteJson(w,http.StatusCreated,post);err !=nil{
		app.InternalServerError(w,r,err)
		return
	}
}

func (app * application) getPostHandler(w http.ResponseWriter,r *http.Request){
 ctx := r.Context()
 id := chi.URLParam(r,"postId")
 postId,err:= strconv.ParseInt(id,10,64)
 if err != nil{
	app.InternalServerError(w,r,err)
 }
 post, err:= app.store.Posts.GetPostById(ctx,postId)
 if err!=nil{
	switch{
	case errors.Is(err,store.ErrNotFound):
	app.NotFoundResponse(w,r,err)
	default:
	app.InternalServerError(w,r,err)
	}
   return
 }
 if err:= WriteJson(w,http.StatusOK,post);err!=nil{
	app.InternalServerError(w,r,err)
	return
 }
}
