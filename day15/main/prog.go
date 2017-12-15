package main

import (
	"fmt"
	"math/big"
)

func main() {

	var aFactor int64 = 16807
	const bFactor int64 = 48271

	const aStart int64 = 65
	const bStart int64 = 8921

	bigA := big.NewInt(aFactor * aStart)
	bigB := big.NewInt(bFactor * bStart)
	fmt.Println(bigA, bigB)

	// not sure what to do with this number
	var divisor = big.NewInt(2147483647)
	var res big.Int
	fmt.Println(res.Rem(bigA, divisor))
}
