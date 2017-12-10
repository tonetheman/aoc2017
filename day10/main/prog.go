package main

import (
	"fmt"
)

const debug = false

//var skipSize = 0
//var currentPosition = 0

// KnotHash - implement from day10
// pass the input ll
// pass the lengths ilens
// pass the skipSize (initial should be 0)
// pass the current Position (initial should be 0)
// returns a chk (res[0]*res[1])
// returns the skipSize
// returns the currentPosition
func KnotHash(ll []int, ilens []int, skipSize, currentPosition int) (int, int, int) {
	var LEN = len(ll)

	if debug {
		fmt.Println(ll)
		fmt.Println("input lens", ilens)
		// skip size starts at 0
		//skipSize := 0
		fmt.Println("skipSize is set to", skipSize)
		// current position in the input array
		//currentPosition := 0
		fmt.Println("current pos", currentPosition)
	}

	// current spot in the lens array
	currentLensIndex := 0
	if debug {
		fmt.Println("current lens index", currentLensIndex)
	}

	// loop count is for debugging
	loopCount := 0
	for loopCount < len(ilens) {
		if debug {
			fmt.Println("ll at start of loop:", ll)
		}

		// amount we need to reverse
		revThisLen := ilens[currentLensIndex]
		if debug {
			fmt.Println("reverse this amount", revThisLen)
		}

		// make a new list to hold the numbers we need to
		// reverse
		newList := make([]int, revThisLen)
		if debug {
			fmt.Println("new List", newList)
		}

		// move those numbers into the list
		// use a temporary so we do not mess up currentPositon (yet)
		tmpCurrentPosition := currentPosition
		for i := 0; i < revThisLen; i++ {
			newList[i] = ll[tmpCurrentPosition]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}
		if debug {
			fmt.Println("newList filled", newList)
		}

		// reverse the list
		for i, j := 0, len(newList)-1; i < j; i, j = i+1, j-1 {
			newList[i], newList[j] = newList[j], newList[i]
		}
		if debug {
			fmt.Println("new list is now", newList)
		}

		// now put the numbers back into the array
		tmpCurrentPosition = currentPosition
		for i := 0; i < revThisLen; i++ {
			ll[tmpCurrentPosition] = newList[i]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}

		if debug {
			fmt.Println("ll is now this", ll)
		}
		currentPosition = (currentPosition + revThisLen + skipSize) % LEN
		if debug {
			fmt.Println("currentPosition has moved to", currentPosition)
		}
		skipSize++
		if debug {
			fmt.Println("skipSize has increased and is now", skipSize)
		}
		currentLensIndex++
		if debug {
			fmt.Println("current Lens index has increased", currentLensIndex)
			fmt.Println("ll at end of loop", ll)
			fmt.Println("------------------------")
		}

		loopCount++
		//if loopCount == 4 {
		//		break
		//	}

	} // end of for loop

	// final 2 numbers
	if debug {
		fmt.Println(ll[0], ll[1])
		fmt.Println("answer:", ll[0]*ll[1])
		fmt.Println("return skipSize:", skipSize)
		fmt.Println("return currentPosition:", currentPosition)
	}
	return ll[0] * ll[1], skipSize, currentPosition
}

func part1() {
	//ll := []int{0, 1, 2, 3, 4}
	//ilens := []int{3, 4, 1, 5}
	ll := make([]int, 256)
	for i := 0; i < 256; i++ {
		ll[i] = i
	}
	fmt.Println(ll)
	ilens := []int{129, 154, 49, 198, 200, 133, 97, 254, 41, 6, 2, 1, 255, 0, 191, 108}
	skipSize := 0
	currentPosition := 0
	KnotHash(ll, ilens, skipSize, currentPosition)
}

func computeSparseHash(a []int) int {
	total := 0
	for i := 0; i < 16; i++ {
		total ^= a[i]
	}
	return total
}

func part2(s string) {
	inputList := make([]int, 0)
	// take the input as a string and use
	// the ascii values
	for _, runeValue := range s {
		inputList = append(inputList, int(runeValue))
	}
	// append something given in the problem
	inputList = append(inputList, 17)
	inputList = append(inputList, 31)
	inputList = append(inputList, 73)
	inputList = append(inputList, 47)
	inputList = append(inputList, 23)
	if debug {
		fmt.Println("s", s)
		fmt.Println("inputList", inputList)
	}

	// set the initial values
	ll := make([]int, 256)
	for i := 0; i < 256; i++ {
		ll[i] = i
	}

	// this section of code is a round for knothash
	skipSize := 0
	currentPosition := 0
	for i := 0; i < 64; i++ {
		_, skipSize, currentPosition = KnotHash(ll, inputList, skipSize, currentPosition)
	}

	if debug {
		fmt.Println("END OF 64 ROUNDS")
		fmt.Println(ll)
	}

	// compute the sparse hash of the result
	res := sp(ll)
	if debug {
		fmt.Println("res is", res)
	}

	// print it all
	for _, val := range res {
		fmt.Printf("%2x", val)
	}
	fmt.Println()
}

func sp(a []int) []int {
	//a := make([]int, 256)
	//for i := 0; i < 256; i++ {
	//	a[i] = rand.Intn(255)
	//}
	//fmt.Println(a)
	res := make([]int, 16)
	for i := 0; i < 16; i++ {
		//fmt.Println(i*16, (i+1)*16, a[i*16:(i+1)*16])
		res[i] = computeSparseHash(a[i*16 : (i+1)*16])
	}
	//fmt.Println(res)
	return res
}
func main() {
	//part2("AoC 2017")
	part2("129,154,49,198,200,133,97,254,41,6,2,1,255,0,191,108")

	//a := []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
	//fmt.Println(compute_spare_hash(a))
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(a[0:3])
	//fmt.Println(a[3 : 3+3])
	//a := make([]int, 256)
	//for i := 0; i < 256; i++ {
	//	a[i] = rand.Intn(255)
	//}
	//fmt.Println(a)
	//sp(a)
}
