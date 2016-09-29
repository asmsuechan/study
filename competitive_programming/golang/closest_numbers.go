package main

import (
	"bufio"
	"fmt"
	"math"
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

func minDiff(a []int) int {
	min := 10000000
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if math.Abs(float64(a[i]-a[j])) < float64(min) {
				min = int(math.Abs(float64(a[i] - a[j])))
			}
		}
	}
	return min
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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, nextInt())
	}
	nums = quicksort(nums)
	min := 100000
	for i := 0; i < n-1; i++ {
		if math.Abs(float64(nums[i]-nums[i+1])) < float64(min) {
			min = int(math.Abs(float64(nums[i] - nums[i+1])))
		}
	}
	for i := 0; i < n-1; i++ {
		if math.Abs(float64(nums[i]-nums[i+1])) == float64(min) {
			fmt.Printf("%d %d ", nums[i], nums[i+1])
		}
	}
}
