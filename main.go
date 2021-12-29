package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/salasberryfin/accumulate-sdk-go/account"
	"github.com/salasberryfin/accumulate-sdk-go/api"
	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
	"github.com/salasberryfin/accumulate-sdk-go/faucet"
	"github.com/salasberryfin/accumulate-sdk-go/get"
)

// KeyResponse
type KeyResponse struct {
	Label      string `json:"name"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Seed       string `json:"seed"`
	Mnemonic   string `json:"mnemonic"`
}

func sendFaucet(apiClient *api.Client, url string) {
	client := faucet.New(apiClient, faucet.Options{})
	response, err := client.SendFaucet(url)
	data := response.Result.Data
	if err != nil {
		log.Fatal("Failed to send faucet: ", err)
	}

	fmt.Println("Faucet was sent successfully: ")
	fmt.Println("\tTx ID: ", data.TxID)
}

func getByURL(apiClient *api.Client, url string) {
	client := get.New(apiClient)
	object, err := client.FromObject(url)
	if err != nil {
		log.Fatal("Something went wrong when using Get - ", err)
	}
	fmt.Println("Object retrieved from Get: ", object)
}

func getTokenAccount(apiClient *api.Client, url string) {
	// if a new tx/faucet was sent just before this function is called
	// it is needed to wait for a few second until it is 'available'
	//time.Sleep(10 * time.Second)
	client := account.New(apiClient)
	details, err := client.Get(url)
	if err != nil {
		log.Println("Failed to retrieve account details - ", err)
	}
	fmt.Println("Account address: ", details.Result.Data.URL)
	fmt.Println("Account balance: ", details.Result.Data.Balance)
}

func createTokenAccount(apiClient *api.Client) (url string) {
	client := account.New(apiClient)
	account, err := client.Generate()
	if err != nil {
		log.Fatal("Failed to create a new account: ", err)
	}
	jsonData := KeyResponse{}
	json.Unmarshal([]byte(account), &jsonData)
	url = jsonData.Label
	fmt.Println("A new account has been generated with the following details: ")
	fmt.Println("\tAddress: ", jsonData.Label)
	fmt.Println("\tPublicKey: ", url)

	return
}

func listAccounts(apiClient *api.Client) {
	client := account.New(apiClient)
	accounts, err := client.List()
	if err != nil {
		log.Fatal("Something went wrong when listing accounts - ", err)
	}
	fmt.Println("Listing the current accounts: ", accounts)
}

func main() {
	options := accsdk.Options{
		JSONOutput: true,
	}
	sdkSession, err := accsdk.NewSession(options)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Server: ", sdkSession.API.Server)
	//fmt.Println("JSON: ", sdkSession.JSONOutput)

	//url := createTokenAccount()
	listAccounts(sdkSession.API)
}
