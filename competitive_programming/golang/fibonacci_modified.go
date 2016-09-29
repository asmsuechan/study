package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func nextInt64() int64 {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return int64(n)
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
	ti := make([]*big.Int, 21)
	ti[0] = big.NewInt(nextInt64())
	ti[1] = big.NewInt(nextInt64())
	n := nextInt()
	for i := 2; i <= n; i++ {
		ti[i] = new(big.Int).Add(ti[i-2], new(big.Int).Mul(ti[i-1], ti[i-1]))
	}
	fmt.Println(ti[n-1])
}
