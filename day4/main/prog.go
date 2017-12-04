package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	inf, err := os.Open("day4.input")
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	count_valid := 0

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		// split array based on space
		// or tab
		textNums := strings.Split(text, " ")
		if len(textNums) == 0 {
			continue
		}

		m := make(map[string]int)
		for i := 0; i < len(textNums); i++ {
			key := textNums[i]
			_, err := m[key]
			if err == false {
				m[key] = 1
			} else {
				m[key]++
			}
		}

		fmt.Println(m)
		// this is  stupid way to do this
		var valid bool = true
		for _, v := range m {
			if v > 1 {
				valid = false
				break
			} else {
				valid = true
			}
		}

		if valid {
			count_valid++
		}

	}

	fmt.Println("valid", count_valid)
}

func main() {
	part1()
}
