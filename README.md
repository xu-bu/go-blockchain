# Intro:
There is a hardhat project in root folder to test contract and several sub projects.
/connectTest is a project to demonstrate how to use go to connect to ethereum.
/interact is an NFT market backend using gin.
/transaction is a project to demonstrate how to use go to proceed transaction.

# Configuration:
Create .env file according to envTemplate.txt.

Run `go install github.com/ethereum/go-ethereum/cmd/abigen@latest` to install abigen, copy abi and bytecode from remix as content of `interactTest/build/MySmartContract.abi` and `interactTest/build/MySmartContract.bin`. Then run `iabigen --abi=./build/sepolia.abi --bin=./build/sepolia.bin --pkg=api --out=./api/sepolia.go` to generate .go api file.

# Usage:
cd each folder, run `go run .` (since main.go invokes some functions in handler.go, need to compile all the files).

# How to get URI of NFT:
Go to https://nft.storage/files/, upload image, then write a metadata JSON file to describe it. The template can be found in transactionTest/assets. Finally, upload metadata file and copy its IPFS URL as URI.
