package interfaces

import (
	"context"
	"e-wallet-ums/external"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userID int) (*external.WalletResponse, error)
	SendNotification(ctx context.Context, recipient string, templateName string, placeholder map[string]string) error
}
