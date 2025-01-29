package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr,maxIdleTime string,maxOpenConns,maxIdleConns int)(*sql.DB,error){
	db,err := sql.Open("postgres",addr)
	if err != nil{
		return nil,err
	}
	duration,err := time.ParseDuration(maxIdleTime)
	if err!= nil{
		return nil,err
	}
	db.SetConnMaxIdleTime(duration)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	if err = db.PingContext(ctx);err!=nil{
		return nil,err
	}

	return db,nil
}