package app

import (
	"database/sql"
	"golangNorthwindRestApi/helper"
	"golangNorthwindRestApi/product"
	"golangNorthwindRestApi/user"

	_ "golangNorthwindRestApi/docs"

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

var (
	productService product.Service
)

func mapUrls(databaseConnection *sql.DB) {

	var (
		productRepository = product.NewRepository(databaseConnection)
	)
	productService = product.NewService(productRepository)

	router.Use(helper.GetCors().Handler)
	router.Mount("/products", product.MakeHttpHandler(productService))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))
}

func mapUrls_mvc() {

	ginRouter.GET("/ping", user.Ping)
	ginRouter.POST("/users", user.Create)
}
