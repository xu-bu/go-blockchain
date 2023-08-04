package test

import (
	"errors"
	"fmt"
	"log"
	"os"

	"goTest/api"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func getContract() (string,error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	API_URL,contractAddress:= os.Getenv("goerliURL"),os.Getenv("CONTRACT_ADDRESS")
	if API_URL == "" || contractAddress==""{
		return "",errors.New("parse error in getContract of handler.go")
	}
	return "",nil
	// 连接到以太坊节点
    client, err := ethclient.Dial(API_URL)
    if err != nil {
        log.Fatal(err)
    }

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, "<YOUR_ADMIN_ACCOUNT_ADDRESS>")

	//deploying smart contract
	deployedContractAddress, tx, instance, err := DeployApi(auth, client) //api is redirected from api directory from our contract go file
	if err != nil {
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex()) // print deployed contract address
}
