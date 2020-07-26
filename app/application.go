package app

import (
	"fmt"
	"golangNorthwindRestApi/database"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var (
	router = chi.NewRouter()
)
var (
	ginRouter = gin.Default()
)

func StartApplication() {
	fmt.Println("application")
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	mapUrls(databaseConnection)

	fmt.Println("about to start the application...")

	http.ListenAndServe(":3000", router)
}

func StarApplicationMVC() {
	fmt.Println("application MVC")
	mapUrls_mvc()

	http.ListenAndServe(":3000", ginRouter)
}
