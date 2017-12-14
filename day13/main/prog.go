package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SCANNER_DOWN = 9
const SCANNER_UP = 8

type layer struct {
	layer_num int
	depth     int
}

func readFile(filename string) []layer {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()
	layers := make([]layer, 0)
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		res := strings.Split(text, ": ")
		layerNumber, _ := strconv.Atoi(res[0])
		depth, _ := strconv.Atoi(res[1])
		layers = append(layers, layer{layerNumber, depth})
	}
	return layers
}

func findMaxScanner(layers []layer) int {
	maxValue := -1
	for _, val := range layers {
		if val.layer_num > maxValue {
			maxValue = val.layer_num
		}
	}
	return maxValue
}

func insert_scanners(ll [][]int) {
	for i := 0; i < len(ll); i++ {
		target := ll[i]
		if len(target) > 0 {
			target[0] = SCANNER_DOWN
		}
	}
}

func findScannerIndex(a []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == SCANNER_DOWN || a[i] == SCANNER_UP {
			return i
		}
	}
	return -1
}

func moveAllScanners(ll [][]int) {
	for i := 0; i < len(ll); i++ {
		if len(ll[i]) > 0 {
			currentIndex := findScannerIndex(ll[i])
			if ll[i][currentIndex] == SCANNER_DOWN {
				newIndex := currentIndex + 1
				if newIndex == len(ll[i]) {
					// change to scanner up
					ll[i][currentIndex] = 0
					// not sure about this yet -2
					ll[i][newIndex-2] = SCANNER_UP
				} else {
					ll[i][currentIndex] = 0
					ll[i][newIndex] = SCANNER_DOWN
				}
			} else {
				// scanner up case
				newIndex := currentIndex - 1
				if newIndex == -1 {
					// change scanner to up
				} else {
					// just move it
					ll[i][currentIndex] = 0
					ll[i][newIndex] = SCANNER_UP
				}
			}
		}
	}
}

func sim(ll [][]int) {
	fmt.Println("sim started...")
	insert_scanners(ll)
	fmt.Println("inserting scanners (9)...")
	fmt.Println(ll)
	fmt.Println("starting loop...")
	count := 0
	playerPosition := -1 //player starts off the grid
	cost := 0
	for {
		fmt.Println("SECOND", count)
		// move player and see if he is caught
		playerPosition++
		fmt.Println("player position is", playerPosition)
		fmt.Println(ll)

		// if the layer has anything in it
		if len(ll[playerPosition]) > 0 {

			// check if the layer has a scanner
			if ll[playerPosition][0] == SCANNER_DOWN ||
				ll[playerPosition][0] == SCANNER_UP {
				// if caught you incur depth * range
				cost += (playerPosition) * len(ll[playerPosition])
				fmt.Println("cost added", cost)

			}
		}

		// now move all the scanners
		fmt.Println("calling move all scanners...")
		moveAllScanners(ll)
		fmt.Println(ll)
		fmt.Println("---------------")
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("end sim")
}

func main() {
	layers := readFile("test.input")
	fmt.Println(layers)
	maxLayerNum := findMaxScanner(layers)
	fmt.Println("max layer num", maxLayerNum)
	ll := make([][]int, maxLayerNum+1)
	for i := 0; i < len(layers); i++ {
		target := layers[i]
		ll[target.layer_num] = make([]int, target.depth)
	}
	fmt.Println(ll)

	sim(ll)
}
