package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

// generate an API key in INFURA and use its endpoints
var infuraURL = "https://mainnet.infura.io/v3/6a9c093bf1fa4c989425eb0276560a1a"
var ganacheURL="http://127.0.0.1:8545"

func main() {
	// connect to infura or local testnet
	// context.Background() is an empty context
	// client,err:=ethclient.DialContext(context.Background(),infuraURL)
	client,err:=ethclient.DialContext(context.Background(),ganacheURL)
	if err!=nil{
		log.Fatalf("Err to create a client:%v",err)
	}
	defer client.Close()
	// get the last block of ethereum blockchain, which is corresponding to the number in etherscan
	block,err:=client.BlockByNumber(context.Background(),nil)
	if err!=nil{
		log.Fatalf("Err to get block:%v",err)
	}
	fmt.Println(block.Number())
}
