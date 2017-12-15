package main

import (
	"fmt"
	"math/big"
)

var aFactor = big.NewInt(16807)
var bFactor = big.NewInt(48271)
var divisor = big.NewInt(2147483647)

func gen(prev *big.Int, factor *big.Int) big.Int {
	var res1 big.Int
	res1.Mul(factor, prev)
	//fmt.Println("tmp mul is", res1.String())
	var res2 big.Int
	res2.Rem(&res1, divisor)
	return res2
}

func main() {

	var aInitialPrev = big.NewInt(65)
	var bInitialPrev = big.NewInt(8921)
	count := 0
	for {
		resA := gen(aInitialPrev, aFactor)
		resB := gen(bInitialPrev, bFactor)
		fmt.Println(resA, resB)

		bytesA := resA.Bytes()
		for _, val := range bytesA {
			fmt.Printf("%b", val)
		}

		// move values over
		aInitialPrev.Set(&resA)
		bInitialPrev.Set(&resB)
		fmt.Println("----------------")
		count++
		if count == 1 {
			break
		}
	}
}
