package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Randomizer(min, max int) {
	c := max - min + 1
	b := make([]byte, c)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	randomNumber := int(nBig.Int64()) + min
	TestCode = randomNumber
}
