package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	name   string
	weight int
	leaf   bool
	kids   []string
}

var p map[string]Program

func readInput() {
	inf, err := os.Open("test.input")
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		// split by space first
		tokens := strings.Split(text, " ")
		// weight is token one
		weightTokens := strings.Split(tokens[1], ")")
		strWeight := weightTokens[0][1:]
		w, _ := strconv.Atoi(strWeight)
		// if there are more than 2 tokens
		// we have a program with children
		isLeaf := len(tokens) == 2
		//fmt.Println(len(tokens), tokens)
		childrens := make([]string, 0)
		if len(tokens) != 2 {
			for i := 3; i < len(tokens); i++ {
				s := tokens[i]
				// peel off the ending , if present
				if strings.HasSuffix(s, ",") {
					childrens = append(childrens, strings.TrimRight(s, ","))
				} else {
					childrens = append(childrens, s)
				}
			}
		}
		//fmt.Println(childrens)
		//fmt.Println("name:", tokens[0])
		//fmt.Println("weightTokens", weightTokens)
		//fmt.Println("strWeight", strWeight)
		//fmt.Println("--------------")
		// token[0] is the name
		p[tokens[0]] = Program{tokens[0], w, isLeaf, childrens}
	}
}

func main() {
	p = make(map[string]Program)
	readInput()
	fmt.Println("done with input")
	fmt.Println(p)
}
