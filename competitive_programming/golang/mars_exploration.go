package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	return sc.Text()
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
	s := nextString()
	str := strings.Split(s, "")
	numberOfChanged := 0
	for i := 0; i < len(str)-2; i += 3 {
		if i%3 == 0 {
			if str[i] != "S" {
				numberOfChanged++
			}
			if str[i+1] != "O" {
				numberOfChanged++
			}
			if str[i+2] != "S" {
				numberOfChanged++
			}
		}
	}
	fmt.Println(numberOfChanged)

}
