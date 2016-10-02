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
	t := nextInt()
	n := make([]int, t)
	prices := make([][]int, t)
	for i := 0; i < t; i++ {
		n[i] = nextInt()
		prices[i] = scanArray(n[i])
	}
	for i := 0; i < t; i++ {
		profit := 0
		max := 0
		for j := n[i] - 1; j >= 0; j-- {
			if prices[i][j] > max {
				max = prices[i][j]
			}
			profit += max - prices[i][j]
		}
		fmt.Println(profit)
	}
}
