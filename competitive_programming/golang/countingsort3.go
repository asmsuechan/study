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

func nextString() string {
	sc.Scan()
	s := sc.Text()
	return s
}

func scanArray(n int) []int {
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, nextInt())
	}
	return ar
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
	x := []int{}
	ar := make([]int, 100)
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		_ = nextString()
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < n; j++ {
			if x[j] <= i {
				ar[i]++
			}
		}
	}
	printArray(ar)
}
