package interfaces

import (
	"context"
	"e-wallet-ums/internal/models"
)

type ILoginService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
}
