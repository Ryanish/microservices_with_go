package model

import model "github.com/ryanish/microservices_with_go/movieapp/metadata/pkg"

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata`
}
