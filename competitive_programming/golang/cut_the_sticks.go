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

func smallest(s []int) int {
	min := 1001
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// cut and create a new array
func cut(s []int) []int {
	min := smallest(s)
	cutSticks := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] != min {
			cutSticks = append(cutSticks, s[i]-min)
		}
	}
	return cutSticks
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	sticks := []int{}
	for i := 0; i < n; i++ {
		sticks = append(sticks, nextInt())
	}
	fmt.Println(len(sticks))
	cutSticks := cut(sticks)
	for len(cutSticks) > 0 {
		fmt.Println(len(cutSticks))
		cutSticks = cut(cutSticks)
	}
}
