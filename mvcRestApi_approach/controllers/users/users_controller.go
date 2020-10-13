package users

import (
	"golangNorthwindRestApi/mvcRestApi_approach/domain/users"
	"golangNorthwindRestApi/mvcRestApi_approach/services"
	"golangNorthwindRestApi/utils/rest_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

// @Summary Create user
// @Tags User
// @Accept json
// @Produce  json
// @Param request body users.User true "User Data"
// @Success 200 {object} users.User "ok"
// @Router /users [post]
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// @Summary Get user
// @Tags User
// @Accept json
// @Produce  json
// @Success 200 {integer} int "ok"
// @Router /users/{user_id} [get]
func GetUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	result, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(http.StatusBadRequest, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// @Summary update user
// @Tags User
// @Accept json
// @Produce  json
// @Param request body users.User true "User Data"
// @Success 200 {object} users.User "ok"
// @Router /users  [put]
func UpdateUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// @Summary delete user
// @Tags User
// @Accept json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /users/{user_id} [delete]
func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	err := services.UsersService.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// @Summary search by status
// @Tags User
// @Accept json
// @Produce  json
// @Param request body string true "User Data"
// @Success 200 {object} users.Users "ok"
// @Router /users//search [get]
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// @Summary login user
// @Tags User
// @Accept json
// @Produce  json
// @Param request body users.LoginRequest true "User Data"
// @Success 200 {object} users.User "ok"
// @Router /users//login [post]
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
