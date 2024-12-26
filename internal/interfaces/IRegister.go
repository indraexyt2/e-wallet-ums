package interfaces

import (
	"context"
	"e-wallet-ums/internal/models"
)

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
