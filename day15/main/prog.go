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

func pb(b []byte) {
	for index, val := range b {
		fmt.Printf("index: %d: %08b ", index, val)
	}
	fmt.Println()
}

func main() {
	// test case
	//var aInitialPrev = big.NewInt(65)
	//var bInitialPrev = big.NewInt(8921)

	var aInitialPrev = big.NewInt(618)
	var bInitialPrev = big.NewInt(814)
	count := 0
	matchCount := 0
	for {
		resA := gen(aInitialPrev, aFactor)
		resB := gen(bInitialPrev, bFactor)
		//fmt.Println(resA, resB)

		bytesA := resA.Bytes()
		bytesB := resB.Bytes()
		//pb(bytesA)
		//pb(bytesB)
		if bytesA[len(bytesA)-1] == bytesB[len(bytesB)-1] {
			if bytesA[len(bytesA)-2] == bytesB[len(bytesB)-2] {
				matchCount++
			}
		}
		// move values over
		aInitialPrev.Set(&resA)
		bInitialPrev.Set(&resB)
		//fmt.Println("----------------")
		count++
		if count == 40000000 {
			break
		}
	}
	fmt.Println("match count", matchCount)
}
