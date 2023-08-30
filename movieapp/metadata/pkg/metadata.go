package model

//Metadata defines the movie metadata. This structure is going to be used by the callers of the service.

type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"'`
	Description string `json:"description"`
	Director    string `json:"director"`
}
