package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	
	"clientTest/config"
)

func main() {
	config:=config.LoadConfig()
	// infuraURL :=config.InfraEndpoint
	ganacheURL:=config.GanacheURL
	// connect to infura or local testnet
	// context.Background() is an empty context
	// client,err:=ethclient.DialContext(context.Background(),InfraEndpoint)
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
