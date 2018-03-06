package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type holder struct {
	num1, num2 int
}

// returns a bunch of holders
// that come in from the input
func readFile() []holder {
	var res []holder
	res = make([]holder, 0)
	inf, err := os.Open("day24.input")
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		// split array based on space
		// or tab
		textNums := strings.Split(text, "/")
		if len(textNums) == 0 {
			continue
		}
		//fmt.Println(textNums)
		num1, _ := strconv.Atoi(textNums[0])
		num2, _ := strconv.Atoi(textNums[1])
		res = append(res, holder{num1, num2})
	}
	return res
}

func findStarts(h []holder) []holder {
	var res []holder
	res = make([]holder, 0)
	for i := 0; i < len(h); i++ {
		if h[i].num1 == 0 {
			res = append(res, h[i])
		}
	}
	return res
}

func main() {
	res := readFile()
	fmt.Println(res)
	fmt.Println("finding starters...")
	res2 := findStarts(res)
	fmt.Println("res2", res2)
}
