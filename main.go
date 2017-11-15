package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Message struct
type Message struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdat"`
}

func main() {
	app := chi.NewRouter()

	// middleware
	app.Use(middleware.RequestID)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.URLFormat)

	// public routes
	app.Get("/", handlerFn)

	// start server
	defer http.ListenAndServe(":3001", app)
	fmt.Println("[Kurama]: Server running...")
}

func handlerFn(res http.ResponseWriter, req *http.Request) {
	// create a new message
	msg := Message{
		ID:        "f80b342c-f90c-4804-9df1-faeb244ab9b8",
		Message:   "Welcome",
		CreatedAt: time.Now(),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	// %+v prints out all the fields in the struct
	// fmt.Printf("%+v", req)
	json.NewEncoder(res).Encode(msg)
}
