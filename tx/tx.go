package tx

import (
	"fmt"
	"log"

	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
)

// GetTransactionDetails returns details for a given tx
func (t Transaction) GetTransactionDetails(hash string) string {
	fmt.Println("Retrieving details for transaction: ", hash)
	details, err := cmd.GetTX(hash)
	if err != nil {
		log.Println("Something went wrong when querying for tx details: ", err)
	}

	return details
}

// CreateTransaction builds a new tx with given source and destination
func (t Transaction) CreateTransaction(from, to, amount string) {

}
