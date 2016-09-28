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
	a := scanArray(n)
	numberOfNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		numberOfNumbers[a[i]]++
	}
	//printArray(numberOfNumbers[:100])
	for i := 0; i < 100; i++ {
		for numberOfNumbers[i] != 0 {
			fmt.Printf("%d ", i)
			numberOfNumbers[i]--
		}
	}
}
