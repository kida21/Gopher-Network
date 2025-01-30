package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kida21/gopher/internal/store"
)
type application struct{
   config config
   store store.Storage
  
}
type dbConfig struct{
    addr string
	maxOpenConns int
    maxIdleConns int
	maxIdleTime string
}

type config struct {
	Addr string
	db dbConfig
	env string
}

func(app *application) mount() *chi.Mux{
  r:= chi.NewRouter()
 r.Use(middleware.RequestID)
 r.Use(middleware.RealIP)
 r.Use(middleware.Logger)
 r.Use(middleware.Recoverer)

  r.Route("/v1",func (r chi.Router)  {
	r.Get("/health",app.HealthCheckHanlder)
	r.Route("/posts",func(r chi.Router) {
		r.Post("/",app.createPostHandler)
	})

  },
 )
  //r.Use(middleware.Timeout(60 * time.Second))
return r
}

func (app *application) run(mux *chi.Mux) error{
	srv := &http.Server{
		Addr: app.config.Addr,
		Handler:mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		IdleTimeout: time.Minute,
	}
	log.Print("starting server on :",app.config.Addr,app.config.env)
    return srv.ListenAndServe()
}