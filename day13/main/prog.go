package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
	"time"
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
	//fmt.Println("\tstep is called for", ll.layerNumber, ll.scannerPos)
	if ll.depth == 0 {
		//fmt.Println("\treturning layer is empty")
		return
	}

	if ll.direction == down {
		//fmt.Println("\tdir is down for", ll.layerNumber)
		if ll.scannerPos == ll.almostDoneIndex {
			//fmt.Println("\tscannerPos is at almost done", ll.layerNumber)
			ll.direction = up
			ll.data[ll.scannerPos] = 0
			ll.scannerPos = ll.depth - 2
			ll.data[ll.scannerPos] = 9
		} else {
			//fmt.Println("\telse case")
			ll.data[ll.scannerPos] = 0
			//fmt.Println("\telse case before scannerPos++", ll.scannerPos)
			ll.scannerPos = ll.scannerPos + 1
			//fmt.Println("\telse case after scannerPos++", ll.scannerPos)
			ll.data[ll.scannerPos] = 9
		}
	} else if ll.direction == up {
		//fmt.Println("\tdir is up for", ll.layerNumber)
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

func findMaxLayerNum(layers []Layer) int {
	maxLayerNum := 0
	for _, ll := range layers {
		if ll.layerNumber > maxLayerNum {
			maxLayerNum = ll.layerNumber
		}
	}
	return maxLayerNum
}

func findALayer(layers []Layer, layerNum int) (Layer, error) {
	for _, ll := range layers {
		if ll.layerNumber == layerNum {
			return ll, nil
		}
	}
	return NewLayer(0, 0), errors.New("invalid layernumber")
}

func rPart2(delay int, layers []Layer, maxLoop int, f *os.File) (bool, int) {
	trace.Start(f)
	person := -1
	gotCaught := false
	totalCost := 0

	for i := 0; i < maxLoop; i++ {
		if delay == 0 {
			person++
			ll := layers[person]
			//fmt.Println("person moved and is now in layer", person, ll)
			//if len(ll.data) > 0 {
			if ll.depth > 0 { // part2 optimization
				if ll.data[0] == 9 {
					//fmt.Println("COST")
					gotCaught = true

					// added this for part2
					// to try to cut time
					// anytime you get caught then
					// you can just stop
					return true, 0

					totalCost += ll.cost()
				}
			}
		} else {
			delay--
		}

		//fmt.Println("stepping layers now...")
		for index := range layers {
			//fmt.Println("\tstep", index)
			// must do this with orig array if
			// you want step to work
			layers[index].step()
		}

		//for _, ll := range layers {
		//fmt.Println("\t", ll)
		//}
		//fmt.Println("--------")
	}
	trace.Stop()
	return gotCaught, totalCost
}

/*
func part2(delay int) (bool, int) {
	//inputLayers := readInputfile("part1.input")
	inputLayers := readInputfile("test.input")
	maxLayerNum := findMaxLayerNum(inputLayers)
	layers := make([]Layer, 0)
	for i := 0; i < maxLayerNum+1; i++ {
		l, err := findALayer(inputLayers, i)
		if err != nil {
			layers = append(layers, NewLayer(i, 0))
		} else {
			layers = append(layers, l)
		}
	}

	// now do the work
	//for i := 0; i < 2; i++ {

	// clear the layers
	for index := range layers {
		layers[index].reset()
	}
	//fmt.Println("before rPArt2 layers", layers)

	//maxLoop := delay + 89 // this is a hard code
	maxLoop := delay + 7 // this is a hard code for test.input
	// delay + MAGIC NUMBER
	//maxLoop := delay + 89 // this is a hard code for part1.input
	gotCaught, cost := rPart2(delay, layers, maxLoop)

	//fmt.Println(delay, gotCaught, cost)
	return gotCaught, cost
	//}

}


func part2Run() {
	for i := 0; i < 11; i++ {
		gotCaught, cost := part2(i)
		if !gotCaught {
			fmt.Println("found it", i, cost)
		}
	}
}
*/

func part2Efficient(delay int, layers []Layer, f *os.File) (bool, int) {

	// clear the layers
	for index := range layers {
		layers[index].reset()
	}
	//fmt.Println("before rPArt2 layers", layers)

	//maxLoop := delay + 89 // this is a hard code
	//maxLoop := delay + 7 // this is a hard code for test.input
	// delay + MAGIC NUMBER
	maxLoop := delay + 89 // this is a hard code for part1.input
	gotCaught, cost := rPart2(delay, layers, maxLoop, f)

	//fmt.Println(delay, gotCaught, cost)
	return gotCaught, cost

}

func part2ERun() {

	f, err := os.Create(time.Now().Format("2006-01-02T150405.pprof"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//trace.Start(f)
	inputLayers := readInputfile("part1.input")
	//inputLayers := readInputfile("test.input")
	maxLayerNum := findMaxLayerNum(inputLayers)
	layers := make([]Layer, 0)
	for i := 0; i < maxLayerNum+1; i++ {
		l, err := findALayer(inputLayers, i)
		if err != nil {
			layers = append(layers, NewLayer(i, 0))
		} else {
			layers = append(layers, l)
		}
	}
	for i := 3000000; i < 4000000; i++ {
		if i%1000 == 0 {
			fmt.Printf("%d - *\n", i)
		}
		gotCaught, cost := part2Efficient(i, layers, f)
		if !gotCaught {
			fmt.Println("found it", i, cost)
		}
	}

	part2Efficient(0, layers, f)
	//trace.Stop()

}

func main() {
	part2ERun()
}
