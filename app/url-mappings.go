package app

import (
	"database/sql"
	"golangNorthwindRestApi/auth/http"
	"golangNorthwindRestApi/auth/repository/db"
	"golangNorthwindRestApi/auth/repository/rest"
	"golangNorthwindRestApi/auth/services/access_token"
	_ "golangNorthwindRestApi/docs"
	"golangNorthwindRestApi/endPointRestApi_approach/product"
	"golangNorthwindRestApi/helper"
	"golangNorthwindRestApi/mvcRestApi_approach/controllers/ping"
	"golangNorthwindRestApi/mvcRestApi_approach/controllers/users"

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

	// users
	ginRouter.GET("/ping", ping.Ping)
	ginRouter.POST("/users", users.Create)
	ginRouter.GET("/users/:user_id", users.GetUser)
	ginRouter.PUT("/users/:user_id", users.UpdateUser)
	ginRouter.DELETE("/users/:user_id", users.DeleteUser)
	ginRouter.GET("/internal/users/search", users.Search)
	ginRouter.POST("/users/login", users.Login)

	//auth
	authHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))
	ginRouter.GET("/oauth/access_token/:access_token_id", authHandler.GetById)
	ginRouter.POST("/oauth/access_token", authHandler.Create)
}
