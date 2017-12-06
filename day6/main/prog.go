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

func cvtInSliceToString(a []int) string {
	txt := []string{}
	for _, val := range a {
		ts := strconv.Itoa(val)
		txt = append(txt, ts)
	}
	return strings.Join(txt, "")
}

func part1() int {

	// this is a slice of ints
	//buckets := getInput("0\t2\t7\t0")
	buckets := getInput("2	8	8	5	4	2	3	1	5	5	1	2	15	13	5	14")

	var seen map[string]int
	seen = make(map[string]int)

	seen[cvtInSliceToString(buckets)] = 0
	seenIndex := 1
	fmt.Println("seen", seen)
	count := 0
	for {
		distOp(buckets)
		ts := cvtInSliceToString(buckets)
		_, ok := seen[ts]
		if ok == false {
			seen[ts] = seenIndex
			seenIndex++
		} else {
			fmt.Println("DUP")
			fmt.Println("current seenIndex", seenIndex)
			fmt.Println("from map", seen[ts])
			fmt.Println("DIST", (seenIndex - seen[ts]))
			count++
			break
		}
		fmt.Println("------------")
		count++
		//if count == 5 {
		//	break
		//}
	}
	fmt.Println("final count", count)
	return count
}

func main() {
	part1()
}
