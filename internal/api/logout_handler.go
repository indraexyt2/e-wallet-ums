package api

import (
	"e-wallet-ums/constants"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c, token)
	if err != nil {
		log.Error("failed to logout: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadGateway,
			false,
			constants.ErrServerError,
			nil,
		)
		c.Abort()
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		constants.SuccessMessage,
		nil,
	)
}
