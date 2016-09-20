package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func calclate(x, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(10, 5))
	fmt.Println(calclate(10, 5))
}
