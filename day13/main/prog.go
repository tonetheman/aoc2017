package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	layers := readFile("test.input")
	fmt.Println(layers)
	maxLayerNum := findMaxScanner(layers)
	fmt.Println("max layer num", maxLayerNum)
}
