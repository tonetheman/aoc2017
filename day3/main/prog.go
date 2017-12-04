package main

import "fmt"

type Layer struct {
	layer          int
	right_bottom   int
	nums_in        int
	right_side_len int
}

func main() {
	var layers []Layer
	layer := 1
	i := 3
	for {

		fmt.Println("layer", layer, "i", i, "ii", i*i)
		layers = append(layers, Layer{layer, i * i, 0, ((layer + 1) * 2) - 1})
		if i*i > 325489 {
			break
		}
		layer++
		i += 2
	}

	// computer the number of numbers in each layer
	for i := 0; i < len(layers); i++ {
		if i != 0 {
			//fmt.Println("fixing nums in", i)
			layers[i].nums_in = layers[i].right_bottom - layers[i-1].right_bottom
		}
	}
	// hard code the first layer
	layers[0].nums_in = 8

	fmt.Println(layers[len(layers)-1])

}
