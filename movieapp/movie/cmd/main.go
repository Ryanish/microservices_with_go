package main

import (
	"github.com/ryanish/microservices_with_go/movieapp/movie/internal/controller/movie"
	metadatagateway "github.com/ryanish/microservices_with_go/movieapp/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/ryanish/microservices_with_go/movieapp/movie/internal/gateway/rating/http"
	httphandler "github.com/ryanish/microservices_with_go/movieapp/movie/internal/handler/http"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
