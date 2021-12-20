package main

import (
	"log"

	accsdk "github.com/salasberryfin/accumulate-sdk-go/client"
)

const accumulateURL = "https://testnet.accumulatenetwork.io/v1"

func main() {
	sdkSession, err := accsdk.MakeSession(accumulateURL)
	if err != nil {
		log.Fatal(err)
	}

	account, err := sdkSession.GenerateAccount()
	if err != nil {
		log.Println("Failed to create a new account: ", err)
	}
	log.Println("New account is: ", account)

	details, err := sdkSession.GetAccount(myTestWallet)
	if err != nil {
		log.Println("Failed to retrieve account details: ", err)
	}
	log.Println(details)
}
