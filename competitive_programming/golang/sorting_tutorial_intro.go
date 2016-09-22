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
	v := nextInt()
	n := nextInt()
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, nextInt())
	}
	for i := 0; i < n; i++ {
		if ar[i] == v {
			fmt.Println(i)
		}
	}
}
