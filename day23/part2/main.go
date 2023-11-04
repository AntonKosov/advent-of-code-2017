package main

import (
	"fmt"
	"math/big"
)

func main() {
	nonPrime := 0
	for i := int64(0); i <= 1_000; i++ {
		if !big.NewInt(105_700 + i*17).ProbablyPrime(0) {
			nonPrime++
		}
	}

	fmt.Printf("Answer: %v\n", nonPrime)
}
