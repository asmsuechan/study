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
	ar := []int{}
	ans := []int{}
	for i := 0; i < t; i++ {
		m := nextInt()
		n := nextInt()
		ar = scanArray(n)
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if ar[j]+ar[k] == m {
					ans = append(ans, j+1)
					ans = append(ans, k+1)
				}
			}
		}
	}
	for i := 0; i < len(ans); i += 2 {
		fmt.Printf("%d %d\n", ans[i], ans[i+1])
	}
}
