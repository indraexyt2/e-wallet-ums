package interfaces

import (
	"context"
	"e-wallet-ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
}
