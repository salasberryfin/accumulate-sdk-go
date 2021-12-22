package account

import (
	"context"
	"fmt"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
	"github.com/AccumulateNetwork/accumulate/types"
	acmeapi "github.com/AccumulateNetwork/accumulate/types/api"
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
func (a Account) GetAccount(url string) (string, error) {
	var res acmeapi.APIDataResponse

	fmt.Println("Request to server: ", a.session.API.Server)
	params := acmeapi.APIRequestURL{}
	params.URL = types.String(url)

	if err := a.session.API.Request(context.Background(), "token-account", params, &res); err != nil {
		return cmd.PrintJsonRpcError(err)
	}

	return cmd.PrintQueryResponse(&res)
}
