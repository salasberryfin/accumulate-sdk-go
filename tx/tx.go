package tx

import (
	"log"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
	"github.com/salasberryfin/accumulate-sdk-go/api"
)

// Options is implemented for potential extra parameters that may be passed
type Options struct {
}

// Transaction is an instance for an Accumulate tx
type Transaction struct {
	API     *api.Client
	options Options
}

// New creates a new instance of a Transaction
func New(apiClient *api.Client) (t *Transaction) {
	t = &Transaction{
		API: apiClient,
	}

	return
}

/*
	TODO: implement get & create tx
	- "token-tx"
	- "token-tx-create"
*/

// Get returns details for a given tx
func (t Transaction) Get(hash string) string {
	// Accumulate API call: "token-tx"
	details, err := cmd.GetTX(hash)
	if err != nil {
		log.Println("Something went wrong when querying for tx details: ", err)
	}

	return details
}

// CreateTransaction builds a new tx with given source and destination
func (t Transaction) CreateTransaction(from, to, amount string) {
	// Accumulate API call: "token-tx-create"
}
