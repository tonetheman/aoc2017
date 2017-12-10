package main

import (
	"fmt"
)

func knothash(ll []int, ilens []int) int {
	var LEN int = len(ll)

	fmt.Println(ll)
	fmt.Println("input lens", ilens)

	// skip size starts at 0
	skipSize := 0
	fmt.Println("skipSize is set to", skipSize)

	// current position in the input array
	currentPosition := 0
	fmt.Println("current pos", currentPosition)

	// current spot in the lens array
	currentLensIndex := 0
	fmt.Println("current lens index", currentLensIndex)

	// loop count is for debugging
	loopCount := 0
	for loopCount < len(ilens) {
		fmt.Println("ll at start of loop:", ll)

		// amount we need to reverse
		revThisLen := ilens[currentLensIndex]
		fmt.Println("reverse this amount", revThisLen)

		// make a new list to hold the numbers we need to
		// reverse
		newList := make([]int, revThisLen)
		fmt.Println("new List", newList)

		// move those numbers into the list
		// use a temporary so we do not mess up currentPositon (yet)
		tmpCurrentPosition := currentPosition
		for i := 0; i < revThisLen; i++ {
			newList[i] = ll[tmpCurrentPosition]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}
		fmt.Println("newList filled", newList)

		// reverse the list
		for i, j := 0, len(newList)-1; i < j; i, j = i+1, j-1 {
			newList[i], newList[j] = newList[j], newList[i]
		}
		fmt.Println("new list is now", newList)

		// now put the numbers back into the array
		tmpCurrentPosition = currentPosition
		for i := 0; i < revThisLen; i++ {
			ll[tmpCurrentPosition] = newList[i]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}

		fmt.Println("ll is now this", ll)
		currentPosition = (currentPosition + revThisLen + skipSize) % LEN
		fmt.Println("currentPosition has moved to", currentPosition)
		skipSize++
		fmt.Println("skipSize has increased and is now", skipSize)
		currentLensIndex++
		fmt.Println("current Lens index has increased", currentLensIndex)

		fmt.Println("ll at end of loop", ll)
		fmt.Println("------------------------")
		loopCount++
		//if loopCount == 4 {
		//		break
		//	}

	} // end of for loop

	// final 2 numbers
	fmt.Println(ll[0], ll[1])
	fmt.Println("answer:", ll[0]*ll[1])
	return ll[0] * ll[1]
}

func main() {
	//ll := []int{0, 1, 2, 3, 4}
	//ilens := []int{3, 4, 1, 5}
	ll := make([]int, 256)
	for i := 0; i < 256; i++ {
		ll[i] = i
	}
	fmt.Println(ll)
	ilens := []int{129, 154, 49, 198, 200, 133, 97, 254, 41, 6, 2, 1, 255, 0, 191, 108}
	knothash(ll, ilens)
}
