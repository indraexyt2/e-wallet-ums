package services

import (
	"context"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo interfaces.IRegisterRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	var (
		log = helpers.Logger
	)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash")
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.RegisterRepo.InsertNewUser(ctx, &request)
	if err != nil {
		log.Error("failed to insert new user")
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil
}
