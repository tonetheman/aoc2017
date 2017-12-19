package main

import "fmt"

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

	// current spot in the lens array
	currentLensIndex := 0

	// loop count is for debugging
	loopCount := 0
	for loopCount < len(ilens) {

		// amount we need to reverse
		revThisLen := ilens[currentLensIndex]

		// make a new list to hold the numbers we need to
		// reverse
		newList := make([]int, revThisLen)

		// move those numbers into the list
		// use a temporary so we do not mess up currentPositon (yet)
		tmpCurrentPosition := currentPosition
		for i := 0; i < revThisLen; i++ {
			newList[i] = ll[tmpCurrentPosition]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}

		// reverse the list
		for i, j := 0, len(newList)-1; i < j; i, j = i+1, j-1 {
			newList[i], newList[j] = newList[j], newList[i]
		}

		// now put the numbers back into the array
		tmpCurrentPosition = currentPosition
		for i := 0; i < revThisLen; i++ {
			ll[tmpCurrentPosition] = newList[i]
			tmpCurrentPosition = (tmpCurrentPosition + 1) % LEN
		}

		currentPosition = (currentPosition + revThisLen + skipSize) % LEN
		skipSize++
		currentLensIndex++

		loopCount++
		//if loopCount == 4 {
		//		break
		//	}

	} // end of for loop

	// final 2 numbers
	return ll[0] * ll[1], skipSize, currentPosition
}

func computeSparseHash(a []int) int {
	total := 0
	for i := 0; i < 16; i++ {
		total ^= a[i]
	}
	return total
}

func fullSparseHash(a []int) []int {
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

func knotHashRound(s string) []int {
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

	// compute the sparse hash of the result
	res := fullSparseHash(ll)
	printSparseHash(res)
	return res
}

func printSparseHash(res []int) {
	// print it all
	for _, val := range res {
		fmt.Printf("%2x", val)
	}
	fmt.Println()
}

// making sure binary is working
// like i think duh
func tCv(s string) string {
	ts := ""
	for _, i := range s {
		if i == 53 {
			ts = ts + "0101"
		}
	}
	return ts
}

func computeUsed(s string) int {

	totalUsed := 0

	for i := 0; i < 128; i++ {
		sKey := fmt.Sprintf("%s-%d", s, i)

		res := knotHashRound(sKey)
		bitString := ""
		for i := 0; i < len(res); i++ {
			ts := fmt.Sprintf("%08b", res[i])
			bitString += ts
		}

		// count the 1s 0==48, 1==49 runes
		for _, val := range bitString {
			totalUsed += (int(val) - 48)
		}

	}

	return totalUsed
}

func computeGrid(s string) int {

	totalUsed := 0

	// create a 128x128 grid
	rows := make([][]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = make([]int, 128)
	}

	for i := 0; i < 128; i++ {
		sKey := fmt.Sprintf("%s-%d", s, i)

		res := knotHashRound(sKey)
		bitString := ""
		for i := 0; i < len(res); i++ {
			ts := fmt.Sprintf("%08b", res[i])
			bitString += ts
		}

		// count the 1s 0==48, 1==49 runes
		for _, val := range bitString {
			totalUsed += (int(val) - 48)
		}

	}

	return totalUsed
}

func messingAround() {
	// checking that i know how to use sprintf
	//sKey := "flqrgnkx"
	//for i := 0; i < 128; i++ {
	//	newKey := fmt.Sprintf("%s-%d", sKey, i)
	//	fmt.Println(newKey)
	//}

	sKey := "flqrgnkx-0"
	res := knotHashRound(sKey)
	fmt.Println("res from part2", res)

	i := res[0]
	fmt.Printf("d %d h %2x b %b\n", i, i, i)
	fmt.Printf("%b\n", 212)
	fmt.Printf("%08b\n", 85) // this is a good format

}

// convert a string of ascii 0,1 into an array of ints
// single dimension only
func sTob(s string) []int {
	b := make([]int,len(s))
	for idx := range s {
		fmt.Println(s[idx])
		if s[idx]==48 {
			b[idx] = 0
		} else if s[idx]==49 {
			b[idx] = 1
		}
	}
	return b
}


func testcase() {
	// test case
	//8108
	fmt.Println("answer:", computeUsed("flqrgnkx"))
}

func main() {
	fmt.Println("answer:", computeUsed("xlqgujun"))
}
