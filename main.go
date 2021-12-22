package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/salasberryfin/accumulate-sdk-go/account"
	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
	"github.com/salasberryfin/accumulate-sdk-go/faucet"
)

const accumulateURL = "https://testnet.accumulatenetwork.io/v1"

type KeyResponse struct {
	Label      string `json:"name"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Seed       string `json:"seed"`
	Mnemonic   string `json:"mnemonic"`
}

func main() {
	options := accsdk.Options{
		JSONOutput: true,
	}
	sdkSession, err := accsdk.NewSession(options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server: ", sdkSession.API.Server)
	fmt.Println("JSON: ", sdkSession.JSONOutput)

	accountClient := account.New(sdkSession)
	account, err := accountClient.GenerateAccount()
	if err != nil {
		log.Println("Failed to create a new account: ", err)
	}
	jsonData := KeyResponse{}
	json.Unmarshal([]byte(account), &jsonData)
	log.Println("New account is: ", account)
	log.Println("New account address: ", jsonData.Label)

	faucetClient := faucet.New(sdkSession)
	faucetResult, err := faucetClient.SendFaucet(jsonData.Label)
	if err != nil {
		log.Println("Failed to send faucet: ", err)
	}
	log.Println("faucetResult: ", faucetResult)

	// Wait for faucet to be processed -> validate test
	time.Sleep(10 * time.Second)
	details, err := accountClient.GetAccount(jsonData.Label)
	if err != nil {
		log.Println("Failed to retrieve account details: ", err)
	}
	log.Println(details)
}
