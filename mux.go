package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/juntakoman123/go_todo_app/clock"
	"github.com/juntakoman123/go_todo_app/config"
	"github.com/juntakoman123/go_todo_app/handler"
	"github.com/juntakoman123/go_todo_app/service"
	"github.com/juntakoman123/go_todo_app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status" : "ok"}`))
	})

	v := validator.New()

	db, cleanup, err := store.New(ctx, cfg)

	if err != nil {
		return nil, cleanup, err
	}

	r := &store.Repository{Clocker: &clock.RealClocker{}}

	at := &handler.AddTask{Service: &service.AddTask{DB: db, Repo: r}, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Service: &service.ListTask{DB: db, Repo: r}}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux, cleanup, nil
}
