package config

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	InfraEndpoint string
	PrivateKey string
	GanacheURL string
	GoerliURL string
	CONTRACT_ADDRESS string
	ACCOUNT1 string
	ACCOUNT2 string
}

func LoadConfig() Config{
	err:=godotenv.Load("../../.env")
	if err!=nil{
		fmt.Println("Error loading .env file")
	}
	return Config{
		InfraEndpoint:os.Getenv("infraEndpoint"),
		PrivateKey:os.Getenv("privateKey"),
		GanacheURL:os.Getenv("ganacheURL"),
		GoerliURL:os.Getenv("goerliURL"),
		CONTRACT_ADDRESS:os.Getenv("CONTRACT_ADDRESS"),
		ACCOUNT1:os.Getenv("ACCOUNT1"),
		ACCOUNT2:os.Getenv("ACCOUNT2"),
	}
}