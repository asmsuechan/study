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
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, nextInt())
	}
	sortedNums := []int{}
	for i := -10000; i <= 10000; i++ {
		for j := 0; j < n; j++ {
			if nums[j] == i {
				sortedNums = append(sortedNums, nums[j])
			}
		}
	}
	fmt.Println(sortedNums[n/2])
}
