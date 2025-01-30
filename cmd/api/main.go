package main

import (
	"log"

	"github.com/kida21/gopher/internal/db"
	"github.com/kida21/gopher/internal/env"
	"github.com/kida21/gopher/internal/store"

	 
)
const version = "0.0.1"
func main() {

	cfg := config{ 
		Addr: env.GetString("ADDR"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR"),
			maxOpenConns: env.GetInt("DB_MAXOPEN_CONNS"),
			maxIdleConns: env.GetInt("DB_MAXIDLE_CONNS"),
			maxIdleTime: env.GetString("DB_MAXIDLE_TIME"),

		},
		env: env.GetString("ENV"),
	}
	db,err := db.New(cfg.db.addr,cfg.db.maxIdleTime,cfg.db.maxOpenConns,cfg.db.maxIdleConns)
	if err != nil{
		log.Panic(err)
	}
	defer db.Close()
	log.Println("database connection pool created")
	store := store.NewStorage(db);
	app := &application{
      config:cfg,
	  store:store,
	}
    
	r := app.mount()
	log.Fatal(app.run(r))

}