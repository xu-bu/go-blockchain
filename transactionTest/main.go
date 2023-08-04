package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	// "crypto/x509"

	"encoding/hex"
	// "encoding/pem"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var sepoliaURL = "https://sepolia.infura.io/v3/6a9c093bf1fa4c989425eb0276560a1a"

// 从account1地址转amount数量ETH到account2地址，需要在.env中填写account1的私钥和API key
func main() {
	// connect to infura or local testnet
	// context.Background() is an empty context
	client, err := ethclient.DialContext(context.Background(), sepoliaURL)
	// client,err:=ethclient.DialContext(context.Background(),ganacheURL)
	if err != nil {
		log.Fatalf("Err to create a client:%v", err)
	}
	defer client.Close()

	account1 := common.HexToAddress("0x52906abb6B9d358eEF7D903cf1ecb521643297f4")
	account2 := common.HexToAddress("0x6c9d7692808c49A65b70401DE0A8a92D6fBd22c7")
	getBalance(client, account1)
	getBalance(client, account2)

	nonce1 := getNonce(client, account1)
	amount := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	// generally, we set gas limit to 21000 when transfer in Ethereum
	// gas limit * gasPrice will be how much Weis we pay for miner
	tx := types.NewTransaction(nonce1, account2, amount, 21000, gasPrice, nil)
	// we need to sign it with our private key to make it valid
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	privateKey,err := privateKeyFromString(os.Getenv("privateKey1"))
	if err != nil {
		log.Println("Parse error!")
		log.Fatal(err)
	}
	// get the id of our test network
	chainID,err:=client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx,err=types.SignTx(tx,types.NewEIP155Signer(chainID),privateKey)
	if err != nil {
		log.Fatal("Sign error!")
		log.Fatal(err)
	}
	err=client.SendTransaction(context.Background(),tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("transaction sent: %s ",tx.Hash().Hex())
	fmt.Println()
	getBalance(client, account1)
	getBalance(client, account2)
}

func privateKeyFromString(privateKeyString string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %v", err)
	}

	privateKey, err := cryptoToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key: %v", err)
	}

	return privateKey, nil
}

func cryptoToECDSA(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.Curve = curve
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKeyBytes)

	if privateKey.PublicKey.X == nil || privateKey.PublicKey.Y == nil {
		return nil, fmt.Errorf("invalid private key")
	}

	return privateKey, nil
}


func getNonce(client *ethclient.Client, addr common.Address) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Fatalln(err)
	}
	return nonce
}

func getBalance(client *ethclient.Client, addr common.Address) {
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatalf("Err to get balance:%v", err)
	}
	// convert unit
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("balance of address",addr," is ",value)
}
