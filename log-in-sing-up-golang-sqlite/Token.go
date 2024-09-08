package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Token int64

func Tokenizator() {
	min := int64(100000000000000000000)
	max := int64(999999999999999999999)

	rangeSize := max - min

	maxBig := big.NewInt(rangeSize + 1)

	randNum, err := rand.Int(rand.Reader, maxBig)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Token = randNum.Int64() + min
}
