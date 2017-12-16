package main

import (
	"fmt"
	"math/big"
)

var aFactor = big.NewInt(16807)
var bFactor = big.NewInt(48271)
var divisor = big.NewInt(2147483647)

var four = big.NewInt(4)
var zero = big.NewInt(0)
var eight = big.NewInt(8)

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

	// part1
	var aInitialPrev = big.NewInt(618)
	var bInitialPrev = big.NewInt(814)
	count := 0      // debug counter
	matchCount := 0 // this is the problem requirement
	for {

		var resA big.Int
		// generate for A first
		for {
			// generate first
			resA = gen(aInitialPrev, aFactor)
			// check that we are divisible evenly by 4
			var chk4 big.Int
			// this does the resA % 4
			chk4.Mod(&resA, four)
			if chk4.Cmp(zero) == 0 {
				// this is good
				break
			} else {
				// keep going
				aInitialPrev.Set(&resA)
			}
		}

		var resB big.Int
		for {
			resB = gen(bInitialPrev, bFactor)
			var chk8 big.Int
			chk8.Mod(&resB, eight)
			if chk8.Cmp(zero) == 0 {
				break
			} else {
				bInitialPrev.Set(&resB)
			}
		}
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
		if count == 5000000 {
			break
		}
	}
	fmt.Println("match count", matchCount)
}
