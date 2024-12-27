package interfaces

import (
	"context"
	"e-wallet-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
}

type ILoginHandler interface {
	Login(c *gin.Context)
}
