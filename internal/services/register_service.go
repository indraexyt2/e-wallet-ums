package services

import (
	"context"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo   interfaces.IUserRepository
	ExternalWallet interfaces.IExtWallet
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.RegisterRepo.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	_, err = s.ExternalWallet.CreateWallet(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil
}
