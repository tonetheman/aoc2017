package main

import "fmt"

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

func main() {
	i := 1
	//const target int = 325489
	const target int = 23
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
	}
	if target >= leftMid(i) && target < bottomMid(i) {
		fmt.Println("between left and bottom")
		layer := i
		bottomLimit := leftMid(i)
		topLimit := bottomMid(i)
		fmt.Println("layer is", layer)
		fmt.Println(target, "sammiched here", topLimit, bottomLimit)

		// now figure horizontal distance to 1
		hDist := target - bottomLimit
		fmt.Println("hDist is", hDist)

		fmt.Println("m dist", layer+hDist)

	}
	if target >= topMid(i) && target < leftMid(i) {
		fmt.Println("not here")
	}
	if target >= rightMid(i) && target < topMid(i) {
		fmt.Println("not here either")
	}

}
