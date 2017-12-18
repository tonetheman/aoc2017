package main

import "fmt"

type Layer struct {
	a int
	b int
}

func (ll *Layer) step() {
	fmt.Println("\tstep is getting called!")
	ll.a = ll.a + 1
}

func main() {
	layers := make([]Layer, 0)
	layers = append(layers, Layer{1, 1})
	layers = append(layers, Layer{2, 2})
	fmt.Println("initial", layers)
	for _, ll := range layers {
		// Tony notes
		// yup ll is a copy of the real value...
		// so this does not work
		fmt.Printf("type of ll %T\n", ll)
		ll.step()
	}
	fmt.Println("WTF: not changed?: step1", layers)

	fmt.Println("looping by hand cause fuck range", layers)
	for i := 0; i < len(layers); i++ {
		layers[i].step()
	}
	fmt.Println("after looping by hand", layers)

	// another way to range
	fmt.Println("before range 2", layers)
	for idx := range layers {
		layers[idx].step()
	}
	fmt.Println("after range 2", layers)

	fakeLayer := Layer{3, 3}
	fmt.Printf("type of fakeLayer %T\n", fakeLayer)
	fmt.Println("fakeLayer", fakeLayer)
	fakeLayer.step()
	fmt.Println("step1:fakeLayer", fakeLayer)
}
