package main

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

// NewLayer - return a new _layer based on what you pass
func NewLayer(layerNumber int, depth int) Layer {
	var tmp Layer
	tmp.layerNumber = layerNumber
	tmp.depth = depth
	tmp.direction = down
	tmp.data = make([]int, depth)
	tmp.scannerPos = 0
	tmp.almostDoneIndex = depth - 1
	return tmp
}

func main() {

}
