package main

import (
	"fmt"
	"math"
)

func rightCorner(n int) int {
	return (2*n - 1) * (2*n - 1)
}

func bottomMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	//fmt.Println("bm", rc, n)
	return rc - (n - 1)
}

func rightMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*7
}

func topMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*5
}

func leftMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*3
}

func part1(target int) int {
	i := 1
	for {
		fmt.Println(i, rightCorner(i), bottomMid(i), leftMid(i), topMid(i), rightMid(i))
		if target < rightCorner(i) {
			fmt.Println("stopping...")
			break
		}
		i++
	}

	// NEED TO CHECK THIS STUFF AND MAKE SURE IT IS RIGHT
	if target >= bottomMid(i) && target < rightCorner(i) {
		fmt.Println("between right corner and bottom")
		fmt.Println("target:", target, bottomMid(i), rightCorner(i))
		aC := math.Abs(float64(target - rightCorner(i)))
		aB := math.Abs(float64(target - bottomMid(i)))
		fmt.Println("ac and ab", aC, aB)
		layer := i - 1
		if aC < aB {
			// closer to the corner
			fmt.Println(layer, aC, layer+int(aC))
		} else {
			fmt.Println(layer, aB, layer+int(aB))
		}
	}
	if target >= leftMid(i) && target < bottomMid(i) {
		fmt.Println("between leftMid and bottomMid")
		fmt.Println("target:", target, leftMid(i), bottomMid(i))
		aC := math.Abs(float64(target - leftMid(i)))
		aB := math.Abs(float64(target - bottomMid(i)))
		fmt.Println("ac and ab", aC, aB)
		layer := i - 1
		if aC < aB {
			fmt.Println(layer, aC, layer+int(aC))
			return layer + int(aC)
		} else {
			fmt.Println(layer, aB, layer+int(aB))
			return layer + int(aB)
		}

	}
	if target >= topMid(i) && target < leftMid(i) {
		fmt.Println("between topMid and leftMid")
		fmt.Println("target:", target, topMid(i), leftMid(i))
		aC := math.Abs(float64(target - topMid(i)))
		aB := math.Abs(float64(target - leftMid(i)))
		fmt.Println("ac and ab", aC, aB)
		layer := i - 1
		if aC < aB {
			// closer to the corner
			fmt.Println(layer, aC, layer+int(aC))
			return layer + int(aC)
		} else {
			fmt.Println(layer, aB, layer+int(aB))
			return layer + int(aB)
		}
	}
	if target >= rightMid(i) && target < topMid(i) {
		fmt.Println("between rightMid and topMid")
		fmt.Println("target:", target, rightMid(i), topMid(i))
		aC := math.Abs(float64(target - rightMid(i)))
		aB := math.Abs(float64(target - topMid(i)))
		fmt.Println("ac and ab", aC, aB)
		layer := i - 1
		if aC < aB {
			fmt.Println(layer, aC, layer+int(aC))
			return layer + int(aC)
		} else {
			fmt.Println(layer, aB, layer+int(aB))
			return layer + int(aB)
		}

	}

	return -1
}

func main() {
	// passed
	//part1(23)
	//fmt.Println("RES", part1(12))
	//fmt.Println("RES", part1(1024))
	//fmt.Println("RES", part1(1025))
	fmt.Println("RES", part1(325489))
}
