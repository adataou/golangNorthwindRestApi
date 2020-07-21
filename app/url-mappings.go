package app

import (
	"golangNorthwindRestApi/product"

	"github.com/go-chi/chi"
)

func mapUrls() {
	var (
		productService product.Service
	)

	router := chi.NewRouter()

	router.Mount("/products", product.MakeHttpHandler(productService))
}
