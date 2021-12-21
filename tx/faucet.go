package tx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AccumulateNetwork/accumulate/types"
	"github.com/salasberryfin/accumulate-sdk-go/config"
)

//Faucet allows to send ACME tokens to a given url for testing purposes
func Faucet(url string) (string, error) {
	var res acmeapi.APIDataResponse
	params := acmeapi.APIRequestURL{}
	params.URL = types.String(url)

	fmt.Println("Fauce request to: ", config.Client.Server)
	if err := config.Client.Request(context.Background(), "faucet", params, &res); err != nil {
		return config.PrintJsonRpcError(err)
	}
	ar := config.ActionResponse{}
	err := json.Unmarshal(*res.Data, &ar)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling create adi result")
	}
	return ar.Print()
}
