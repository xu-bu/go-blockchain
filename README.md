# Intro:

There is a hardhat project to test NFT contract.

In /transactionTest folder it demonstrates how to use Go to interact with contract in both local and devnet env.
/interactTest is an NFT market backend using gin.

# Set up:

```
go mod tidy
yarn
```

Create .env file according to .env.template.

Copy abi and bytecode from remix as content of `interactTest/build/MySmartContract.abi` and `interactTest/build/MySmartContract.bin`.

Then run

`abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go`

to generate .go api file.

# Usage:

cd each folder, run `go run .` (since main.go invokes some functions in handler.go, need to compile all the files).

# How to get URI of NFT:

Go to https://nft.storage/files/, upload image, then write a metadata JSON file to describe it. The template can be found in transactionTest/assets. Finally, upload metadata file and copy its IPFS URL as URI.
