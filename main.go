package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// new chi router
	router := chi.NewRouter()
	// add middleware logger
	router.Use(middleware.Logger)

	// map the path /hello to basicHandler
	router.Get("/hello", basicHandler)

	// creates a http server on address localhost:3000 and uses a http handler called router which is a chi http router
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// start the server and serve the http requests
	err := server.ListenAndServe()

	// in case unable to start the server, print the error message
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

// handler which has two parameters http request, http response writer
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
