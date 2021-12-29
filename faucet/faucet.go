package faucet

import (
	"github.com/salasberryfin/accumulate-sdk-go/api"
)

// Options is implemented to allow passing extra parameters if needed
type Options struct {
}

// Faucet is an instance for an Accumulate faucet tx
type Faucet struct {
	APIClient    *api.Client
	extraOptions Options
}

// New creates a new instance of a Faucet
func New(apiClient *api.Client, options Options) *Faucet {
	f := Faucet{
		APIClient:    apiClient,
		extraOptions: options,
	}

	return &f
}

// SendFaucet sends test tokens
func (f Faucet) SendFaucet(url string) (resp api.FaucetResponse, err error) {
	req := api.GenericRequest{
		JSONRpc: "2.0",
		ID:      0,
		Method:  "faucet",
		Params: api.Params{
			URL: url,
		},
	}

	err = f.APIClient.SendRequestV1(req, &resp)

	return
}
