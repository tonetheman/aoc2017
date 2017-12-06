package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput(s_input string) []int {
	var buckets []int
	buckets = make([]int, 0)
	s_data := strings.Split(s_input, "\t")
	for _, sVal := range s_data {
		iVal, _ := strconv.Atoi(sVal)
		buckets = append(buckets, iVal)
	}
	//fmt.Println(buckets)
	return buckets
}

func distOp(buckets []int) {
	var maxIndex int = -1
	var maxValue int = -1
	for _index, val := range buckets {
		if val > maxValue {
			maxValue = val
			maxIndex = _index
		}
	}
	fmt.Println("dist maxValue and index", maxValue, maxIndex)
	blocksToDist := maxValue
	buckets[maxIndex] = 0 // take from bucket
	// now pus them all back
	for {
		maxIndex = (maxIndex + 1) % len(buckets)
		buckets[maxIndex] = buckets[maxIndex] + 1
		blocksToDist--
		if blocksToDist == 0 {
			break
		}
	}
	fmt.Println("AFTER DIST", buckets)
}

func checkForMatch(orig []int, a []int) bool {
	for index, val := range orig {
		if a[index] != val {
			return false
		}
	}
	return true
}

func part1() {

	buckets := getInput("0\t2\t7\t0")
	// Need to save off buckets

	orig := make([]int, len(buckets))
	copy(orig, buckets)
	count := 0
	for {
		distOp(buckets)
		fmt.Println(count, "match",
			checkForMatch(orig, buckets))
		fmt.Println("------------")
		count++
		if count == 5 {
			break
		}
	}
}

func main() {

}
