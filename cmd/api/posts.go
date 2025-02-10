package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
    "github.com/kida21/gopher/internal/store"
)
type postKey string
const postCtx postKey = "post"
type CreatePayLoad struct {
	Title string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required,max=1000"`
	Tags []string `json:"tags" validate:"required"`
}

func (app *application) createPostHandler(w http.ResponseWriter,r *http.Request) {
   var payload CreatePayLoad
   if err := ReadJson(w,r,&payload);err != nil{
	app.BadRequestResponse(w,r,err)
   }
   ctx := r.Context()
    if err:= Validate.Struct(payload);err!=nil{
	app.BadRequestResponse(w,r,err)
	return
}
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
post:=getPostFromCtx(r)
 if err:= WriteJson(w,http.StatusOK,post);err!=nil{
	app.InternalServerError(w,r,err)
	return
 }
}

func(app *application) deletePostHandler(w http.ResponseWriter,r *http.Request){
   ctx := r.Context()
 id := chi.URLParam(r,"postId")
 postId,err:= strconv.ParseInt(id,10,64)
 if err != nil{
	app.InternalServerError(w,r,err)
 }
 if err:=app.store.Posts.Delete(ctx,postId);err!=nil{
	switch{
	case errors.Is(err,store.ErrNotFound):
	app.NotFoundResponse(w,r,err)
	default:
	app.InternalServerError(w,r,err)
	}
 }
}
type UpdatePayload struct{
	Content *string `json:"content" validate:"omitempty,max=100"`
	Title *string `json:"title" validate:"omitempty,max=1000"`
	
}
func(app *application) updatePostHandler(w http.ResponseWriter,r *http.Request){
 post:=getPostFromCtx(r)
 
 var payload UpdatePayload
    if err:= ReadJson(w,r,&payload);err !=nil{
		app.BadRequestResponse(w,r,err)
		return
	}
if err:= Validate.Struct(payload);err!=nil{
	app.BadRequestResponse(w,r,err)
	return
}
if payload.Content!=nil{
	post.Content = *payload.Content
}
if payload.Title!=nil{
	post.Title=*payload.Title
}

if err:=app.store.Posts.Update(r.Context(),post);err!=nil{
	app.InternalServerError(w,r,err)
	return
 }
 if err:= WriteJson(w,http.StatusOK,post);err!=nil{
	app.InternalServerError(w,r,err)
	return
 }

}

func(app *application) postsContextMiddleware(next http.Handler) http.Handler{
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
//  if err:= WriteJson(w,http.StatusOK,post);err!=nil{
// 	app.InternalServerError(w,r,err)
// 	return
//  }
 ctx = context.WithValue(ctx,postCtx,post)
 next.ServeHTTP(w,r.WithContext(ctx))
 })
 
}
func getPostFromCtx(r*http.Request)(*store.Post){
  post,_:= r.Context().Value(postCtx).(*store.Post)
  return post
}