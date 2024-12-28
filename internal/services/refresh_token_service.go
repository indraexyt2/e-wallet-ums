package services

import (
	"context"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"github.com/pkg/errors"
	"time"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, claimToken *helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, claimToken.UserID, claimToken.Username, claimToken.FullName, "token", time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	err = s.UserRepo.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update token")
	}

	resp.Token = token
	return resp, nil
}
