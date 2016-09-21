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

func numberOfA(s string) int {
	number := 0
	cutS := strings.Split(s, "")
	for _, r := range cutS {
		if r == "a" {
			number += 1
		}
	}
	return number
}

func main() {
	var s string
	fmt.Scan(&s)
	sc.Split(bufio.ScanWords)
	n := nextInt()

	adr := n % len(s)
	fmt.Println(n/len(s)*numberOfA(s) + numberOfA(s[0:adr]))
}
