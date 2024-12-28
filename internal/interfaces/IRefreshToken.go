package interfaces

import (
	"context"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type IRefreshToken interface {
	RefreshToken(ctx context.Context, refreshToken string, claimToken *helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}
