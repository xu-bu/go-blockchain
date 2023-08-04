package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type ResponseData struct {
	Message string `json:"message"`
}

func testRes() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Print("Hello from the server!")

		// 在页面中输出字符串
		// w.Header().Set("Content-Type", "text/plain")
		// w.WriteHeader(http.StatusOK)
		// fmt.Fprint(w, "Response from the server!")

		ytData := ResponseData{
			Message: "Response from the server!",
		}

		// Convert the data to JSON format
		jsonData, err := json.Marshal(ytData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to indicate JSON data
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data to the response writer
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

func getETHBalance() {

	// connect to infura or local testnet
	// context.Background() is an empty context
	client, err := ethclient.DialContext(context.Background(), goerliURL)
	// client,err:=ethclient.DialContext(context.Background(),ganacheURL)
	if err != nil {
		log.Fatalf("Err to create a client:%v", err)
	}
	defer client.Close()

	// go to etherscan choose a random addr and get its balance here
	addr := common.HexToAddress("0x52906abb6B9d358eEF7D903cf1ecb521643297f4")
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatalf("Err to get balance:%v", err)
	}
	// convert unit
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(value)
}

func getABI(){

}

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

    // 合约ABI
    // contractABI := `[{"constant":false,"inputs":[],"name":"yourFunctionName","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
    // abiJSON, err := abi.JSON(strings.NewReader(contractABI))
    // if err != nil {
    //     log.Fatal(err)
    // }

    // 调用函数
    contract := ethclient.NewContract(contractAddress, abiJSON)
    // callData, err := contract.Pack("yourFunctionName")
    // if err != nil {
    //     log.Fatal(err)
    // }

    // msg := ethereum.CallMsg{
    //     To:   &contractAddress,
    //     Data: callData,
    // }

    // result, err := client.CallContract(context.Background(), msg, nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // // 处理返回结果
    // var res string
    // err = contract.UnpackIntoInterface(&res, "yourFunctionName", result)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println("Function result:", res)
}
