package product

import (
	"net/http"

	"github.com/go-chi/chi"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	return r
}
