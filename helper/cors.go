package helper

import (
	"github.com/go-chi/cors"
)

func GetCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "UPDATE", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "C-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
