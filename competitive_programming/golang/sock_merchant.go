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
	ar := scanArray(n)
	numberOfPairs := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ar[i] != 0 && ar[i] == ar[j] {
				numberOfPairs++
				ar[j] = 0
				break
			}
		}
	}
	fmt.Println(numberOfPairs)
}
