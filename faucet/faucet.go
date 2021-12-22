package faucet

import (
	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
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
func (f Faucet) SendFaucet(url string) (string, error) {
	result, err := cmd.Faucet(url)
	if err != nil {
		return "", err
	}

	return result, nil
}
