package cmd

import (
	"e-wallet-ums/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (d *Dependency) MiddlewareValidateAuth(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization header is empty")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	if d.UserRepository == nil {
		panic("UserRepository is nil")
	}

	_, err := d.UserRepository.GetUserSessionByToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session by token: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	claim, err := helpers.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("failed to validate token: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}
	c.Set("token", claim)

	c.Next()
	return
}