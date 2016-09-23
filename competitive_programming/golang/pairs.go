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

func quicksort(a []int) []int {
	n := len(a)
	left, right, eq := []int{}, []int{}, []int{}
	if n > 1 {
		eq = append(eq, a[0])
		eqval := eq[0]
		for i := 0; i < n; i++ {
			if a[i] > eqval {
				right = append(right, a[i])
			} else if a[i] < eqval {
				left = append(left, a[i])
			}
		}
		if len(left) > 1 {
			left = quicksort(left)
		}
		if len(right) > 1 {
			right = quicksort(right)
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

func checkDiff(a []int, k int) int {
	n := len(a)
	numberOfDiff := 0
	for i := 0; i < n-1; i++ {
		if a[i+1]-a[i] <= k {
			//find half of element between i and n
			half := (i + n - 1) / 2
			//half < n
			//incremental search
			if a[half]-a[i] > k {
				for j := i + 1; j < half; j++ {
					if a[j]-a[i] == k {
						numberOfDiff++
						break
					}
				}
				//decremental search
			} else {
				for j := n - 1; j >= half; j-- {
					if a[j]-a[i] == k {
						numberOfDiff++
						break
					}
				}
			}
		}
	}
	return numberOfDiff
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, nextInt())
	}
	fmt.Println(checkDiff(quicksort(ar), k))
}
