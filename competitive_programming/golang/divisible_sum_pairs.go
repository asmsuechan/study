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

func main() {
	sc.Split(bufio.ScanWords)
	nk := []int{}
	ai := []int{}
	for i := 0; i < 2; i++ {
		nk = append(nk, nextInt())
	}
	n := nk[0]
	k := nk[1]
	for i := 0; i < n; i++ {
		ai = append(ai, nextInt())
	}

	numberOfDivisible := 0
	for i := 0; i < n; i++ {
		for j := 1; i+j < n; j++ {
			if (ai[i]+ai[i+j])%k == 0 {
				numberOfDivisible += 1
			}
		}
	}
	fmt.Println(numberOfDivisible)
}
