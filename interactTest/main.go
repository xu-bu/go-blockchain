package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/6a9c093bf1fa4c989425eb0276560a1a"
var ganacheURL="http://127.0.0.1:8545"
var sepoliaURL="https://sepolia.infura.io/v3/6a9c093bf1fa4c989425eb0276560a1a"

func main() {
	// connect to infura or local testnet
	// context.Background() is an empty context
	client,err:=ethclient.DialContext(context.Background(),sepoliaURL)
	// client,err:=ethclient.DialContext(context.Background(),ganacheURL)
	if err!=nil{
		log.Fatalf("Err to create a client:%v",err)
	}
	defer client.Close()

	// go to etherscan choose a random addr and get its balance here
	addr:=common.HexToAddress("0x52906abb6B9d358eEF7D903cf1ecb521643297f4")
	balance,err:=client.BalanceAt(context.Background(),addr,nil)
	if err!=nil{
		log.Fatalf("Err to get balance:%v",err)
	}
	// convert unit
	fBalance:=new(big.Float)
	fBalance.SetString(balance.String())
	value:=new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))
	fmt.Println(value)
}
