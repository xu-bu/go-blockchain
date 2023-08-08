package main

import (
	"context"
	"encoding/json"
	_ "errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"goTest/interactTest/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var API_URL string
var contractAddress string
var privateKey string
var client *ethclient.Client
var auth *bind.TransactOpts
var account string

func init() {
	// 获取.env
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Getwd in main of interactTest/test/test.go:", err)
	}

	envPath := wd + "/../.env"
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error getting .env in main of interactTest/test/test.go:", err)
	}

	// 注意给全局变量赋值不能用:=
	API_URL, contractAddress, privateKey, account = os.Getenv("goerliURL"), os.Getenv("CONTRACT_ADDRESS"), os.Getenv("privateKey"), os.Getenv("ACCOUNT1")
	if API_URL == "" || contractAddress == "" || privateKey == "" || account == "" {
		log.Fatal("Error parsing .env in main of interactTest/test/test.go:")
	}

	// 连接到以太坊节点
	client, err = ethclient.Dial(API_URL)
	if err != nil {
		log.Fatal(err, "connection error in deployContract of test.go")
	}
}

func createNFT(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := api.NewApi(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal("new api error in createNFT of test.go\n", err)
	}

    privateKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        log.Fatal("failed to parse private key")
    }

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatal("failed to get chain ID")
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatal("failed to create transaction options")
    }

    auth.From = common.HexToAddress(account)

	tokenURI := "test"
	_, err = conn.CreateNFT(auth, tokenURI) // conn call the balance function of deployed smart contract
	if err != nil {
		log.Fatal("createNFT error in createNFT of test.go\n", err)
	}
	fmt.Println("create NFT successfully.")
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message": "success",
	}
	// Encode the NFT object as JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}

func mintContract(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
}