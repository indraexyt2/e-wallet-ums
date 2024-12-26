package api

import (
	"e-wallet-ums/constants"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

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

	res, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to register new user: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
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
		res,
	)
	return
}
