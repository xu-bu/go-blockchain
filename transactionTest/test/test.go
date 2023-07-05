package main

import (
	"fmt"
	"math/big"

)

func main() {
	amount := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	fmt.Println(amount)
	
}