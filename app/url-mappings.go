package app

import (
	"database/sql"
	"golangNorthwindRestApi/helper"
	"golangNorthwindRestApi/mvcRestApi/controllers/ping"
	"golangNorthwindRestApi/mvcRestApi/controllers/users"
	"golangNorthwindRestApi/endPointRestApi/product"

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

	ginRouter.GET("/ping", ping.Ping)
	ginRouter.POST("/users", users.Create)
	ginRouter.GET("/users/:user_id", users.GetUser)
	ginRouter.PUT("/users/:user_id", users.UpdateUser)
	ginRouter.DELETE("/users/:user_id", users.DeleteUser)
	ginRouter.GET("/internal/users/search", users.Search)
}
