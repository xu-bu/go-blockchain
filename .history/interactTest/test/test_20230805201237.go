package main

import (
	"context"
	"crypto/ecdsa"
	_ "errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"goTest/interactTest/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var API_URL string
var contractAddress string
var privateKey string
var client *ethclient.Client
var auth *bind.TransactOpts
var account string

func main() {
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
	API_URL, contractAddress, privateKey, account = os.Getenv("sepoliaURL"), os.Getenv("CONTRACT_ADDRESS"), os.Getenv("privateKey"), os.Getenv("ACCOUNT1")
	if API_URL == "" || contractAddress == "" || privateKey == "" || account == "" {
		log.Fatal("Error parsing .env in main of interactTest/test/test.go:")
	}

	// 连接到以太坊节点
	client, err = ethclient.Dial(API_URL)
	if err != nil {
		log.Fatal(err, "connection error in deployContract of test.go")
	}

	
	startTest()
}

func startTest() {
	// interact测试
	createNFT()
	log.Println("createNFT test pass")
	getOwnerOfNFT()
	log.Println("getOwnerOfNFT test pass")
}

func getAccountAddress(privateKey string)(string){
	ECDSAPrivateKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        log.Fatal("failed to parse private key in getAccountAddress")
    }

	publicKey := ECDSAPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return accountAddress.Hex()
}

func getOwnerOfNFT() {
	conn, err := api.NewApi(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal("new api error in createNFT of test.go\n", err)
	}
	tokenID:="1"
	accountAddress:=getAccountAddress(privateKey)
	//调用ownOf的时候不是用auth
	opts := &bind.CallOpts{
		From:    common.HexToAddress(accountAddress), // 发送者地址
	}
	bigIntTokenID := new(big.Int)
	bigIntTokenID.SetString(tokenID, 10) // Base 10
	res, err := conn.OwnerOf(opts, bigIntTokenID) // conn call the balance function of deployed smart contract
	if err != nil {
		log.Fatal("createNFT error in createNFT of test.go\n", err)
	}
	log.Println(res)
}


// contract交互测试
func createNFT() {
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
	reply, err := conn.CreateNFT(auth, tokenURI) // conn call the balance function of deployed smart contract
	fmt.Println(reply)
	if err != nil {
		log.Fatal("createNFT error in createNFT of test.go\n", err)
	}
	fmt.Println("create NFT successfully.")
}
