package interfaces

import (
	"context"
	"e-wallet-ums/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
