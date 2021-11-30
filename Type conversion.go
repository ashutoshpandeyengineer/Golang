package main

import (
	"fmt"
)

func main() {
	var x int = 100
	fmt.Printf("%v,%T\n", x, x)

	var y float32
	y = float32(x)
	fmt.Printf("%v,%T\n", y, y)
}
