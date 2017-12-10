package main

import (
	"fmt"
	"sort"
)

func main() {

	ll := [5]int{0, 1, 2, 3, 4}
	fmt.Println(ll)
	ilens := [4]int{3, 4, 1, 5}
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
	for {
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
			tmpCurrentPosition = (tmpCurrentPosition + 1) % 5
		}
		fmt.Println("newList filled", newList)
		// WTF ... seriously there must be an easier way to do
		// this. wut a mess. for sorting go is shit
		sort.Sort(sort.Reverse(sort.IntSlice(newList)))
		fmt.Println("newList sorted", newList)

		// now put the numbers back into the array
		tmpCurrentPosition = currentPosition
		for i := 0; i < revThisLen; i++ {
			ll[tmpCurrentPosition] = newList[i]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % 5
		}

		fmt.Println("ll is now this", ll)
		currentPosition += revThisLen
		fmt.Println("currentPosition has moved to", currentPosition)
		skipSize++
		fmt.Println("skipSize has increased and is now", skipSize)
		currentLensIndex++
		fmt.Println("current Lens index has increased", currentLensIndex)

		fmt.Println("ll at end of loop", ll)
		fmt.Println("------------------------")
		loopCount++
		if loopCount == 2 {
			break
		}

	} // end of for loop

}
