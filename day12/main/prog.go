package main

import (
	"bufio"
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

// WTF golang no built in stack?
type stack struct {
	vec []int
}

func (s stack) empty() bool {
	return len(s.vec) == 0
}
func (s stack) peek() int {
	return s.vec[len(s.vec)-1]
}
func (s stack) push(val int) {
	s.vec = append(s.vec, val)
}
func (s stack) pop() int {
	tmp := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return tmp
}

func main() {
	programs := readFile("test.input")
	fmt.Println(programs[0])

	var alreadyFound map[int]bool
	alreadyFound = make(map[int]bool)
	alreadyFound[0] = true
	var myStack stack
	stack.push(0)
	lookingForDirects := stack.pop()

}
