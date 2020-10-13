package http

import (
	atDomain "golangNorthwindRestApi/auth/domain/access_token"
	"golangNorthwindRestApi/auth/services/access_token"
	"golangNorthwindRestApi/utils/rest_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

//@Summary Get access_token by Id
//@Tags Auth
//@Accept json
//@Produce json
//@Success 200 {sting} string "ok"
//@Router /oauth/access_token/{access_token_id} [get]
func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

//@Summary Post access_token
//@Tags Auth
//@Accept json
//@Produce json
//@Success 200 {string} string "ok"
//@Router /oauth/access_token [post]
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}
