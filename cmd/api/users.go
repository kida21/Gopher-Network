package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kida21/gopher/internal/store"
)

func (app *application) getUserHandler(w http.ResponseWriter,r *http.Request) {
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
   if err:=WriteJson(w,http.StatusOK,user);err!=nil{
	app.InternalServerError(w,r,err)
	return
   }
}