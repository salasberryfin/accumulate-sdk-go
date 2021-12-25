package account

import (
	"fmt"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"

	acmeapi "github.com/salasberryfin/accumulate-sdk-go/api"
	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
)

type Options struct {
}

// Account resource
type Account struct {
	session accsdk.Session
	options Options
}

// New creates a new instance of type Account
func New(session accsdk.Session) *Account {
	a := Account{
		session: session,
	}

	return &a
}

// GenerateAccount allows a SDK Session to create a new Accumulate account
func (a Account) GenerateAccount() (string, error) {
	return cmd.GenerateKey("")
}

// GetAccount retrieves information for the given account identifier
func (a Account) GetAccount(url string) ([]byte, error) {
	fmt.Println("Request to server: ", a.session.API.Server)
	req := acmeapi.GenericRequest{
		JSONRpc: "2.0",
		ID:      0,
		Method:  "token-account",
		Params: acmeapi.Params{
			URL: url,
		},
	}

	resp, err := a.session.API.SubmitRequestV1(&req, true)

	return resp, err
}
