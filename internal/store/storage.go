package store

import (
	"context"
	"database/sql"
	"errors"
)
var (
	ErrNotFound = errors.New("record not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context,*Post) error
		GetPostById(context.Context,int64)(*Post,error)
		Delete(context.Context,int64) error
		Update(context.Context,*Post)(error)
	}
	Users interface{
		Create(context.Context,*User) error
		GetUserById(context.Context,int64)(*User,error)
	}
}

func NewStorage( db *sql.DB) Storage{
	return Storage{
		Posts:&PostStore{db:db} ,
		Users:&UserStore{db:db},
		
	}
}