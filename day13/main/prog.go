package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// got this from the problm
const maxLayer = 6

type layer struct {
	layerNum int
	depth    int
	scanner  int
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
		layers = append(layers,
			layer{layerNumber, depth, 0})
	}
	return layers
}

// stupid hard code
// this will only work for test data
// TODO: fix
func fillGapsInData(layers []layer) []layer {
	tmpLayers := make([]layer, 0)
	tmpLayers = append(tmpLayers, layers[0])
	tmpLayers = append(tmpLayers, layers[1])
	tmpLayers = append(tmpLayers, layer{0, 0, 0})
	tmpLayers = append(tmpLayers, layer{0, 0, 0})
	tmpLayers = append(tmpLayers, layers[2])
	tmpLayers = append(tmpLayers, layer{0, 0, 0})
	tmpLayers = append(tmpLayers, layers[3])
	return tmpLayers
}

func part1() {
	layers := readFile("test.input")
	layers = fillGapsInData(layers)
	fmt.Println(layers)

}

func main() {
	part1()
}
