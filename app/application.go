package app

import (
	"fmt"
	"golangNorthwindRestApi/database"
	"golangNorthwindRestApi/product"

	_ "github.com/go-sql-driver/mysql"
)

func StartApplication() {
	fmt.Println("application")
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository = product.NewRepository(databaseConnection)
	)

	fmt.Println(databaseConnection)
	fmt.Println("finish")
}
