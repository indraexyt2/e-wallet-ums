package external

import (
	"bytes"
	"context"
	"e-wallet-ums/helpers"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type Wallet struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	Balance int `json:"balance"`
}

type ExtWallet struct{}

func (e *ExtWallet) CreateWallet(ctx context.Context, userID int) (*Wallet, error) {
	req := Wallet{UserID: userID}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "localhost") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "/")
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create wallet http request")
	}

	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create wallet")
	}

	result := &Wallet{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	defer resp.Body.Close()
	return result, nil
}
