package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kida21/gopher/internal/store"
)
type userKey string
const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter,r *http.Request) {
   user:= getUserFromCtx(r)
   if err:=WriteJson(w,http.StatusOK,user);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
}
type Payload struct{
	UserId int64 `json:"user_id"`
}

func(app *application) followUserHandler(w http.ResponseWriter,r*http.Request){
  FollowUser:= getUserFromCtx(r)
  var payload Payload
  if err:= ReadJson(w,r,&payload);err!=nil{
	app.BadRequestResponse(w,r,err)
	return
  }
   if err:=app.store.Followers.Follow(r.Context(),FollowUser.ID,payload.UserId);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
   if err:=WriteJson(w,http.StatusNoContent,FollowUser);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
}

func(app *application) unfollowUserHandler(w http.ResponseWriter,r*http.Request){
   unFollowedUser:= getUserFromCtx(r)
  var payload Payload
  if err:= ReadJson(w,r,&payload);err!=nil{
	app.BadRequestResponse(w,r,err)
	return
  }
   if err:=app.store.Followers.Unfollow(r.Context(),unFollowedUser.ID,payload.UserId);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
   if err:=WriteJson(w,http.StatusNoContent,unFollowedUser);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
}

func(app *application) userContextMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId,err:=strconv.ParseInt( chi.URLParam(r,"userId"),10,64)
   ctx:=r.Context();
   if err!=nil{
	app.BadRequestResponse(w,r,err)
	return
   }
  
   user,err:=app.store.Users.GetUserById(ctx,userId)
   if err!=nil{
     switch err{
	case store.ErrNotFound:
	app.NotFoundResponse(w,r,err)
	return
	default:
	app.InternalServerError(w,r,err)
	return
	}
	}
	ctx = context.WithValue(ctx,userCtx,user)
	next.ServeHTTP(w,r.WithContext(ctx))

	})
}

func getUserFromCtx(r*http.Request)(*store.User){
	user,_:= r.Context().Value(userCtx).(*store.User)
	return user
}