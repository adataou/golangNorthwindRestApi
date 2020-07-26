package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Create(c *gin.Context) {
	fmt.Println("user controller works")
	//var user User
}
