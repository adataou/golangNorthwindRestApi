package main

import (
	"golangNorthwindRestApi/app"
)

// @title Northwind API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	app.StartApplication()

	app.StarApplicationMVC()
}
