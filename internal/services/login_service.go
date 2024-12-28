package services

import (
	"context"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	var (
		resp models.LoginResponse
		now  = time.Now()
	)

	userDetail, err := s.UserRepo.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get user detail")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return resp, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.Wrap(err, "failed to insert new user session")
	}

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken
	return resp, nil
}
