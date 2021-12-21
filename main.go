package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
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
	sdkSession, err := accsdk.MakeSession(accumulateURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sdk session initialized: ", sdkSession)

	account, err := sdkSession.GenerateAccount()
	if err != nil {
		log.Println("Failed to create a new account: ", err)
	}
	jsonData := KeyResponse{}
	json.Unmarshal([]byte(account), &jsonData)
	fmt.Println("A new account was generated:")
	fmt.Println("Address: ", jsonData.Label)
	fmt.Println("Private key: ", jsonData.PrivateKey)

	faucetResult, err := sdkSession.Faucet(jsonData.Label)
	if err != nil {
		log.Println("Failed to send faucet: ", err)
	}
	log.Println("faucetResult: ", faucetResult)

	// Wait for faucet to be processed -> validate test
	time.Sleep(10 * time.Second)
	details, err := sdkSession.GetAccount(jsonData.Label)
	if err != nil {
		fmt.Println("Failed to retrieve account details for ", jsonData.Label+": ", err)
	}
	log.Println(details)
}
