package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return n
}

func divide(a []int) []int {
	n := len(a)
	left, right, eq := []int{}, []int{}, []int{}
	if n > 0 {
		eq = append(eq, a[0])
		eqval := eq[0]
		for i := 0; i < n; i++ {
			if a[i] > eqval {
				right = append(right, a[i])
			} else if a[i] < eqval {
				left = append(left, a[i])
			}
		}
	}
	return append(left, append(eq, right...)...)
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println("")
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, nextInt())
	}
	printArray(divide(ar))
}
