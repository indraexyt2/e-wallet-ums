package interfaces

import (
	"context"
	"e-wallet-ums/external"
)

type IExtWallet interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
}
