package account

import (
	"fmt"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"

	"github.com/salasberryfin/accumulate-sdk-go/api"
)

// Options is implemented for potential extra parameters that may be passed
type Options struct {
}

// Account resource
type Account struct {
	APIClient    *api.Client
	extraOptions Options
}

// New creates a new instance of type Account
func New(apiClient *api.Client) (a *Account) {
	a = &Account{
		APIClient: apiClient,
	}

	return a
}

// Generate allows a SDK Session to create a new Accumulate account
func (a Account) Generate() (string, error) {
	return cmd.GenerateKey("")
}

// List all accounts in current workspace
func (a Account) List() (string, error) {
	return cmd.ListAccounts()
}

// Get retrieves information for the given account identifier
func (a Account) Get(url string) (resp api.GetTokenAccountResponse, err error) {
	fmt.Println("Request to server: ", a.APIClient)
	req := api.GenericRequest{
		JSONRpc: "2.0",
		ID:      0,
		Method:  "token-account",
		Params: api.Params{
			URL: url,
		},
	}

	err = a.APIClient.SendRequestV1(req, &resp)

	return
}
