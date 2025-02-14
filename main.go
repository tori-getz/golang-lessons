package main

import (
	"fmt"
	"math"
)

func main() {
	height, kg := 1.67, 53.0
	IMT := kg / math.Pow(height, 2)
	fmt.Print(IMT)
}
