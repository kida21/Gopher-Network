package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter,r*http.Request){

	ctx:= r.Context()
	feed,err:=app.store.Posts.GetUserFeed(ctx,int64(64))
	if err!=nil{
		app.InternalServerError(w,r,err)
		return
	}
	if err:= WriteJson(w,http.StatusOK,feed);err!=nil{
		app.InternalServerError(w,r,err)
		return
	}
}