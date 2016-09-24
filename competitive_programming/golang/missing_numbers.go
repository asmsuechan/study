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
	A := scanArray(n)
	m := nextInt()
	B := scanArray(m)
	ans := []int{}
	var numListA [1000011]int
	var numListB [1000011]int
	for i := 0; i < len(A); i++ {
		numListA[A[i]]++
	}
	for i := 0; i < len(B); i++ {
		numListB[B[i]]++
	}
	for i := 0; i < 1000011; i++ {
		if numListB[i]-numListA[i] != 0 {
			ans = append(ans, i)
		}
	}
	printArray(ans)
}
