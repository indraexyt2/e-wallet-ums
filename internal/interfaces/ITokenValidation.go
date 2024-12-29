package interfaces

import (
	"context"
	pb "e-wallet-ums/cmd/proto/tokenvalidation"
	"e-wallet-ums/helpers"
)

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error)
}
