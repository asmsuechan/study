package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func nextString() string {
	sc.Scan()
	s := sc.Text()
	return s
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
	//s := os.Args[1]
	s := nextString()
	r := regexp.MustCompile(`[A-Z]`)
	str := r.ReplaceAllString(string(s), "")
	fmt.Println(len(s) - len(str) + 1)
}
