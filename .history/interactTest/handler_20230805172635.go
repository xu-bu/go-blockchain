package main

import (
	_ "context"
	_"encoding/json"
	_ "errors"
	_"fmt"
	"log"
	"net/http"
	_ "net/url"
	"os"

	"goTest/interactTest/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_"github.com/julienschmidt/httprouter"
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

func createNFT(c *gin.Context) {
	mii
	init()
	conn, err := api.NewApi(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal("new api error in createNFT of test.go\n", err)
	}

    // privateKey, err := crypto.HexToECDSA(privateKey)
    // if err != nil {
    //     log.Fatal("failed to parse private key")
    // }

    // chainID, err := client.ChainID(context.Background())
    // if err != nil {
    //     log.Fatal("failed to get chain ID")
    // }

    // auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    // if err != nil {
    //     log.Fatal("failed to create transaction options")
    // }

    auth.From = common.HexToAddress(account)

	tokenURI := "test"
	_, err = conn.CreateNFT(auth, tokenURI) // conn call the balance function of deployed smart contract
	if err != nil {
		log.Fatal("createNFT error in createNFT of test.go\n", err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}

func mintContract(c *gin.Context) {
	address,StatusOK := c.GetPostForm("accountAddress")
	if !StatusOK{
		panic("未获取到name参数")
	}
	c.JSON(http.StatusOK,gin.H{
		"address":address,
	})
}






