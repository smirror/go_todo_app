package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"todo_app/clock"
	"todo_app/config"
	"todo_app/handler"
	"todo_app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset-utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}
	r := store.Repository{Clocker: clock.RealClock{}}
	at := &handler.AddTask{DB: db, Repo: r, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{DB: db, Repo: r}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, nil
}
