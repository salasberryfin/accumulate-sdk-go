package api

import (
	"encoding/json"
	"errors"
)

type APIV1Response interface {
	GetError() error
}

// FaucetData
type FaucetData struct {
	TxID      string `json:"txid"`
	Hash      string `json:"hash"`
	CodeSpace string `json:"codespace"`
}

// TokenAccountData
type TokenAccountData struct {
	URL           string `json:"url"`
	TokenURL      string `json:"tokenurl"`
	KeyBookUrl    string `json:"keyBookUrl"`
	Balance       string `json:"balance"`
	TxCount       int    `json:"txCount"`
	Nonce         int    `json:"nonce"`
	CreditBalance string `json:"creditBalance"`
}

type APIResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type APIResultResponse struct {
	Type    string           `json:"type"`
	MdRoot  string           `json:"mdRoot"`
	Data    *json.RawMessage `json:"data"`
	Sponsor string           `json:"sponsor"`
	KeyPage string           `json:"keyPage"`
	TxId    string           `json:"txid"`
}

// TODO: fix this function to return the specific struct for each type
func (r APIResponse) ParseData() (APIResponse, error) {
	res := r.Result
	var dst interface{}
	switch res.Type {
	case "token-account":
		dst = new(TokenAccountData)
	case "":
		dst = new(FaucetData)
	}
	err := json.Unmarshal(*res.Data, dst)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (r APIResponse) GetError() error {
	noError := APIResponseError{}
	if r.Error != noError {
		return errors.New("API request caused an error")
	}

	return nil
}

type APIResponse struct {
	JSONRpc string            `json:"jsonrpc"`
	Result  APIResultResponse `json:"result"`
	Error   APIResponseError  `json:"error"`
	ID      int               `json:"id"`
}
