package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"goTest/interactTest/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var API_URL string
var contractAddress string
var privateKey string
var client *ethclient.Client
var auth *bind.TransactOpts

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Getwd in main of interactTest/test/test.go:", err)
	}

	envPath := wd + "/../.env"
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error getting .env in main of interactTest/test/test.go:", err)
	}

	API_URL, contractAddress, privateKey = os.Getenv("goerliURL"), os.Getenv("CONTRACT_ADDRESS"), os.Getenv("privateKey")
	if API_URL == "" || contractAddress == "" || privateKey == "" {
		log.Fatal("Error parsing .env in main of interactTest/test/test.go:")
	}

	client, err = ethclient.Dial(API_URL)
	if err != nil {
		log.Fatal(err, "connection error in deployContract of test.go")
	}

	// create auth and transaction package for deploying smart contract
	auth = getAccountAuth()

	contractAddress,err=deployContract()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(contractAddress)
}


// 根据私钥算出账户地址然后获得一个auth，之后才能以该用户的名义进行操作
func getAccountAuth() *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("account address:", fromAddress)
	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	// fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(2500000036) // in wei
	return auth
}

// 根据abigen生成的接口interactTest/api/MySmartContract.go进行合约部署
func deployContract() (contractAddress string, e error) {

	//deploying smart contract
	deployedContractAddress, tx, _, err := api.DeployApi(auth, client)
	_ = tx
	if err != nil {
		fmt.Println(err)
		return "", errors.New("deployment error in deployContract of test.go")
	}

	return deployedContractAddress.Hex(), nil
}

