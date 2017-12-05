package main

import "fmt"

type Position struct {
	row, col int
}

var vals [100]int
var mvals map[Position]int

func fill(n int) {
	if n == 0 {
		mvals[Position{0, 0}] = 1
		vals[n] = 1
		return
	}
	if n == 1 {
		mvals[Position{0, 1}] = 1
		vals[n] = 1
		fill(0)
		return
	}
	last := n - 1
	if vals[last] == 0 {
		fill(last)
	}
	// now compute your value

}

func main() {
	mvals = make(map[Position]int)
	fill(2)
	fmt.Println(mvals)
}
