package main

import (
	"fmt"
)

var skipSize = 0
var currentPosition = 0

func knothash(ll []int, ilens []int) int {
	var LEN int = len(ll)

	fmt.Println(ll)
	fmt.Println("input lens", ilens)

	// skip size starts at 0
	//skipSize := 0
	fmt.Println("skipSize is set to", skipSize)

	// current position in the input array
	//currentPosition := 0
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
	fmt.Println("return skipSize:", skipSize)
	fmt.Println("return currentPosition:", currentPosition)
	return ll[0] * ll[1]
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
	knothash(ll, ilens)
}

func compute_spare_hash(a []int) int {
	total := 0
	for i := 0; i < 16; i++ {
		total ^= a[i]
	}
	return total
}

func part2(s string) {
	input_list := make([]int, 0)
	for _, runeValue := range s {
		input_list = append(input_list, int(runeValue))
	}
	input_list = append(input_list, 17)
	input_list = append(input_list, 31)
	input_list = append(input_list, 73)
	input_list = append(input_list, 47)
	input_list = append(input_list, 23)
	fmt.Println("s", s)
	fmt.Println("input_list", input_list)

	ll := make([]int, 256)
	for i := 0; i < 256; i++ {
		ll[i] = i
	}

	for i := 0; i < 64; i++ {
		knothash(ll, input_list)
	}

	fmt.Println("END OF 64 ROUNDS")
	fmt.Println(ll)
	res := sp(ll)
	fmt.Println("res is", res)
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
		fmt.Println(i*16, (i+1)*16, a[i*16:(i+1)*16])
		res[i] = compute_spare_hash(a[i*16 : (i+1)*16])
	}
	fmt.Println(res)
	return res
}
func main() {
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
