package api

import (
	"context"
	pb "e-wallet-ums/cmd/proto/tokenvalidation"
	"e-wallet-ums/constants"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/interfaces"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	pb.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	var (
		token = req.GetToken()
		log   = helpers.Logger
	)

	if token == "" {
		log.Error("token is empty")
		return &pb.TokenResponse{
			Message: "token is empty",
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &pb.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &pb.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.FullName,
			Email:    claimToken.Email,
		},
	}, nil
}
