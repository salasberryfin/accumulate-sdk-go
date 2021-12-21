package accsdk

import (
	"context"

	"github.com/AccumulateNetwork/accumulate/types"
	acmeapi "github.com/AccumulateNetwork/accumulate/types/api"
	"github.com/salasberryfin/accumulate-sdk-go/config"
	"github.com/salasberryfin/accumulate-sdk-go/core"
)

// GenerateAccount allows a SDK Session to create a new Accumulate account
func (s Session) GenerateAccount() (string, error) {
	return core.GenerateKey("")
}

// GetAccount retrieves information for the given account identifier
func (s Session) GetAccount(url string) (string, error) {
	var res acmeapi.APIDataResponse

	params := acmeapi.APIRequestURL{}
	params.URL = types.String(url)

	if err := s.Client.Request(context.Background(), "token-account", params, &res); err != nil {
		return config.PrintJsonRpcError(err)
	}

	return config.PrintQueryResponse(&res)
}

func (s Session) Faucet(url string) (string, error) {
	result, err := core.Faucet(url)
	if err != nil {
		return "", err
	}

	return result, nil
}
