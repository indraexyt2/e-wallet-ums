package api

import (
	"e-wallet-ums/constants"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshToken
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Error("failed to get claim in context")
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

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim to claim token")
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

	resp, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, tokenClaim)
	if err != nil {
		log.Error("failed to refresh token: ", err)
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
		resp,
	)
}
