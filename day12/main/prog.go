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

func isAlreadyFound(af map[int]bool, lookingFor int) bool {
	_, present := af[lookingFor]
	// golang sucks for this type of stuff
	// i think true means it is there
	return present
}

func main() {
	programs := readFile("test.input")

	var programMap map[int]connection
	programMap = make(map[int]connection)
	for _, p := range programs {
		//fmt.Println(p)
		programMap[p.parent] = p
	}
	fmt.Println(programMap)

	// by prob definition 0 is already found
	var alreadyFound map[int]bool
	alreadyFound = make(map[int]bool)
	// put 0 in the already found map
	alreadyFound[0] = true

	// save off 0 into the stack
	var ll lifo
	ll.push(0)

	count := 0
	fmt.Println("starting now....")
	for {
		if ll.empty() {
			fmt.Println("lifo is empty")
			break
		} else {
			fmt.Println("lifo is not empty keep going")
		}
		currentGuy, _ := ll.pop()
		currentProgram := programMap[currentGuy]
		fmt.Println("currentGuy from lifo is", currentGuy, currentProgram)
		for _, childIndex := range currentProgram.children {
			// find out if the child index is directly
			// connected to anyone
			childProgram := programMap[childIndex]
			fmt.Println("child program of currentGuy", childProgram)
			if !isAlreadyFound(alreadyFound, childIndex) {
				ll.push(childIndex)
				alreadyFound[childIndex] = true
			}
			// this is a connected person too check

			for _, innerChildIndex := range childProgram.children {
				fmt.Println("now checking inner child", innerChildIndex)
				if !isAlreadyFound(alreadyFound, innerChildIndex) {
					fmt.Println("not in alreadyFound pushing on lifo and putting in alreadyFound", innerChildIndex)
					ll.push(innerChildIndex)
					alreadyFound[innerChildIndex] = true
				}

			}
		}
		fmt.Println("lifo at the end of loop", ll)
		count++
		if count == 1 {
			break
		}
	} // end of while loop

	fmt.Println("final already found", alreadyFound)
}
