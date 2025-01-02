package api

import (
	"e-wallet-ums/constants"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.LoginRequest
		resp models.LoginResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			constants.ErrFailedBadRequest,
			nil,
		)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			constants.ErrFailedBadRequest,
			nil,
		)
		return
	}

	resp, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("failed on login service: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			constants.ErrServerError,
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		constants.SuccessMessage,
		resp,
	)
	return
}
