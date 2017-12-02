package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open the file
	inf, err := os.Open("part1.input")
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	total := 0

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		// split array based on space
		// or tab
		textNums := strings.Split(text, "\t")
		if len(textNums) == 0 {
			continue
		}
		fmt.Println(len(textNums), textNums)
		fmt.Println("about to make slice this long", len(textNums))
		intRow := make([]int, len(textNums))
		for i := 0; i < len(textNums); i++ {
			intRow[i], _ = strconv.Atoi(textNums[i])
		}
		fmt.Println(intRow)
		fmt.Println("-----------------")
		sort.Ints(intRow)
		diff := intRow[len(intRow)-1] - intRow[0]
		total += diff
	}
	fmt.Println(total)
}
