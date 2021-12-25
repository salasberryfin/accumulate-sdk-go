package faucet

import (
	acmeapi "github.com/salasberryfin/accumulate-sdk-go/api"
	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
)

// Options additional configuration options for the faucet
type Options struct {
}

// Faucet is an instance for an Accumulate tx
type Faucet struct {
	session accsdk.Session
	options Options
}

// New creates a new Faucet instance
func New(s accsdk.Session, options ...Options) *Faucet {
	f := Faucet{
		session: s,
	}

	return &f
}

// SendFaucet creates sends test tokens
func (f Faucet) SendFaucet(url string) (resp []byte, err error) {
	req := acmeapi.GenericRequest{
		JSONRpc: "2.0",
		ID:      0,
		Method:  "faucet",
		Params: acmeapi.Params{
			URL: url,
		},
	}

	resp, err = f.session.API.SubmitRequestV1(&req, true)

	return
}
