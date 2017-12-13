package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type connection struct {
	parent   int
	children []int
}

func readFile(filename string) []connection {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	programs := make([]connection, 0)
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		res := strings.Split(text, " <-> ")
		res2 := strings.Split(res[1], ",")
		parent, _ := strconv.Atoi(res[0])
		children := make([]int, 0)
		for _, v := range res2 {
			tmp, _ := strconv.Atoi(strings.Trim(v, " "))
			children = append(children, tmp)
		}
		c := connection{parent, children}
		programs = append(programs, c)
	}
	return programs
}

type lifo struct {
	data []int
}

// WHAT I HAVE IS A LIFO
func (s *lifo) push(val int) {
	s.data = append(s.data, val)
}
func (s lifo) empty() bool {
	return len(s.data) == 0
}
func (s *lifo) pop() (int, error) {
	if s.empty() {
		// in my case -1 is nothing
		// if this was a real stack
		// then return an error
		return -1, errors.New("empty stack")
	}
	var tmp int
	// this pulls the last thing inserted
	// then resets data to the rest of the list
	tmp, s.data = s.data[0], s.data[1:]
	return tmp, nil
}
func (s lifo) len() int {
	return len(s.data)
}

func main() {
	programs := readFile("test.input")
	fmt.Println(programs[0])

	var alreadyFound map[int]bool
	alreadyFound = make(map[int]bool)
	alreadyFound[0] = true
	var myStack lifo
	myStack.push(0)

	lookingForDirects := myStack.pop()
	fmt.Println(lookingForDirects)

}
