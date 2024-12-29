package services

import (
	"context"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		log        = helpers.Logger
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		log.Error("failed to validate token: ", err)
		return claimToken, err
	}

	_, err = s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to get user session by token")
	}

	return claimToken, nil
}
