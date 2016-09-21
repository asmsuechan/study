package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	a := []int{}
	b := []int{}
	for i := 0; i < 3; i++ {
		a = append(a, scanInt())
	}
	for i := 0; i < 3; i++ {
		b = append(b, scanInt())
	}

	point_a := 0
	point_b := 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			point_a += 1
		} else if b[i] > a[i] {
			point_b += 1
		}
	}
	fmt.Printf("%d %d\n", point_a, point_b)
}
