package main

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"sort" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (f PaginatedFeedQuery) Parse(r *http.Request)(PaginatedFeedQuery,error){
	qs:= r.URL.Query()
	limit:= qs.Get("limit")
	if limit!=""{
	 l,err:= strconv.Atoi(limit)
	 if err!=nil{
		return f,err
	 }
	 f.Limit=l
	}
	offset:= qs.Get("offset")
	if offset!=""{
	 o,err:= strconv.Atoi(offset )
	 if err!=nil{
		return f,err
	 }
	 f.Offset=o
	}
	sort:=qs.Get("sort")
	if sort!=""{
		f.sort=sort
	}
	return f,nil
}