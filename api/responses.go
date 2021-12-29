package api

import (
	"errors"
)

// APIV1Response that must be satisfied by all response types to dynamycally
// parse the JSON responses with the required fields
type APIV1Response interface {
	GetError() error
}

// APIResponseError represents the error format when requests fail
type APIResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// GetError allows to check for error responses when submiting the request
func (r APIV1ResponseEnvelope) GetError() error {
	noError := APIResponseError{}
	if r.Error != noError {
		return errors.New("API request caused an error: " + r.Error.Message)
	}

	return nil
}

// APIV1ResponseEnvelope is the base of all API responses
type APIV1ResponseEnvelope struct {
	_struct struct{}         `codec:",omitempty,omitemptyarray"`
	JSONRpc string           `json:"jsonrpc"`
	Error   APIResponseError `json:"error"`
	ID      int              `json:"id"`
}

// GetTokenAccountResponse is the HTTP response received when retrieving details for an
// account
type GetTokenAccountResponse struct {
	APIV1ResponseEnvelope
	Result TokenAccountResult `json:"result"`
}

// TokenAccountResult is the specific result field
type TokenAccountResult struct {
	Type    string      `json:"type"`
	MdRoot  string      `json:"mdRoot"`
	Data    AccountData `json:"data"`
	Sponsor string      `json:"sponsor"`
	KeyPage string      `json:"keyPage"`
	TxId    string      `json:"txid"`
}

// AccountData is the specific data field
type AccountData struct {
	URL           string `json:"url"`
	TokenURL      string `json:"tokenurl"`
	KeyBookURL    string `json:"keyBookUrl"`
	Balance       string `json:"balance"`
	TxCount       int    `json:"txCount"`
	Nonce         int    `json:"nonce"`
	CreditBalance string `json:"creditBalance"`
}

// FaucetResponse is the HTTP response received when creating a new faucet
type FaucetResponse struct {
	APIV1ResponseEnvelope
	Result FaucetResult
}

// FaucetResult is the specific result field
type FaucetResult struct {
	Type    string     `json:"type"`
	MdRoot  string     `json:"mdRoot"`
	Data    FaucetData `json:"data"`
	Sponsor string     `json:"sponsor"`
	KeyPage string     `json:"keyPage"`
	TxID    string     `json:"txid"`
}

// FaucetData is the specific data field
type FaucetData struct {
	TxID      string `json:"txid"`
	Hash      string `json:"hash"`
	CodeSpace string `json:"codespace"`
}

// GetResponse is the HTTP response received when retriieving an object by URL
type GetResponse struct {
	APIV1ResponseEnvelope
	Result GetResult
}

// GetResult is the specific result field
type GetResult struct {
	Type    string  `json:"type"`
	MdRoot  string  `json:"mdRoot"`
	Data    GetData `json:"data"`
	Sponsor string  `json:"sponsor"`
	KeyPage string  `json:"keyPage"`
}

// GetData is the specific data field
type GetData struct {
	URL           string `json:"url"`
	TokenURL      string `json:"tokenurl"`
	KeyBookURL    string `json:"keyBookUrl"`
	Balance       string `json:"balance"`
	TxCount       int    `json:"txCount"`
	Nonce         int    `json:"nonce"`
	CreditBalance string `json:"creditBalance"`
}

// AdiResponse is the HTTP response received when requesting information by ADI
type AdiResponse struct {
	APIV1ResponseEnvelope
	Result AdiResult
}

// AdiResult is the specific result
type AdiResult struct {
	Type    string  `json:"type"`
	MdRoot  string  `json:"mdRoot"`
	Data    AdiData `json:"data"`
	Sponsor string  `json:"sponsor"`
	KeyPage string  `json:"keyPage"`
	TxID    string  `json:"txid"`
}

// AdiData is the specific data
type AdiData struct {
	URL         string
	PublicKey   string
	KeyBookName string
	KeyPageName string
}

// TokenTxResponse
type TokenTxResponse struct {
	APIV1ResponseEnvelope
	Result TokenTxResult
}
