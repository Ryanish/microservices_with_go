package main

import (
	"github.com/ryanish/microservices_with_go/movieapp/rating/internal/controller/rating"
	httphandler "github.com/ryanish/microservices_with_go/movieapp/rating/internal/handler/http"
	"github.com/ryanish/microservices_with_go/movieapp/rating/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
