package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	down = 1
	up   = 2
)

// Layer holds the info about each scanner layer
type Layer struct {
	layerNumber     int
	depth           int
	direction       int
	data            []int
	scannerPos      int
	almostDoneIndex int
}

func (ll *Layer) reset() {
	if ll.depth != 0 {
		for idx := range ll.data {
			ll.data[idx] = 0
		}
		ll.data[0] = 9
	}
	ll.scannerPos = 0
	ll.direction = down
}
func (ll Layer) cost() int {
	return ll.layerNumber * ll.depth
}
func (ll Layer) repr() {
	fmt.Printf("l %d - ", ll.layerNumber)
	for _, val := range ll.data {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
func (ll *Layer) step() {
	if ll.depth == 0 {
		return
	}
	if ll.direction == down {
		if ll.scannerPos == ll.almostDoneIndex {
			ll.direction = up
			ll.data[ll.scannerPos] = 0
			ll.scannerPos = ll.depth - 2
			ll.data[ll.scannerPos] = 9
		} else {
			ll.data[ll.scannerPos] = 0
			ll.scannerPos++
			ll.data[ll.scannerPos] = 9
		}

	} else if ll.direction == up {
		if ll.scannerPos == 0 {
			ll.direction = down
			ll.data[ll.scannerPos] = 0
			ll.scannerPos = 1
			ll.data[ll.scannerPos] = 9
		} else {
			ll.data[ll.scannerPos] = 0
			ll.scannerPos--
			ll.data[ll.scannerPos] = 9
		}

	}
}

// NewLayer - return a new _layer based on what you pass
func NewLayer(layerNumber int, depth int) Layer {
	var tmp Layer
	tmp.layerNumber = layerNumber
	tmp.depth = depth
	tmp.direction = down
	if depth != 0 {
		tmp.data = make([]int, depth)
		tmp.data[0] = 9
	}
	tmp.scannerPos = 0
	tmp.almostDoneIndex = depth - 1
	return tmp
}

func layerTest() {
	l := NewLayer(0, 4)

	for i := 0; i < 8; i++ {
		l.repr()
		l.step()
	}
}

func readInputfile(filename string) []Layer {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	layers := make([]Layer, 0)
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		// split the text line on the colon
		text := scanner.Text()
		snums := strings.Split(text, ":")

		// convert the text to numbers watch the trims
		// cause go is picky
		num1, _ := strconv.Atoi(strings.Trim(snums[0], ""))
		num2, _ := strconv.Atoi(strings.Trim(snums[1], " "))
		//fmt.Println(snums[0], snums[1], num1, num2)

		// make a new layer for what we just read
		layers = append(layers, NewLayer(num1, num2))
	}
	//fmt.Println(layers)
	return layers
}

func main() {
	//layerTest()
	readInputfile("part1.input")

}
