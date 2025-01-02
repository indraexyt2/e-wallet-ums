package services

import (
	"context"
	"e-wallet-ums/internal/interfaces"
	"e-wallet-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo interfaces.IUserRepository
	External     interfaces.IExternal
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

	_, err = s.External.CreateWallet(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	s.External.SendNotification(ctx, request.Email, "register", map[string]string{
		"full_name": request.FullName,
	})

	resp := request
	resp.Password = ""

	return resp, nil
}
