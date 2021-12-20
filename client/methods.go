package accsdk

import (
	"context"
	"fmt"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
	"github.com/AccumulateNetwork/accumulate/types"
	acmeapi "github.com/AccumulateNetwork/accumulate/types/api"
)

// GenerateAccount allows a SDK Session to create a new Accumulate account
func (s Session) GenerateAccount() (string, error) {
	return cmd.GenerateKey("")
}

// GetAccount retrieves information for the given account identifier
func (s Session) GetAccount(url string) (string, error) {
	var res acmeapi.APIDataResponse

	fmt.Println("Request to server: ", s.API.Server)
	params := acmeapi.APIRequestURL{}
	params.URL = types.String(url)

	if err := s.API.Request(context.Background(), "token-account", params, &res); err != nil {
		return cmd.PrintJsonRpcError(err)
	}

	return cmd.PrintQueryResponse(&res)
}
