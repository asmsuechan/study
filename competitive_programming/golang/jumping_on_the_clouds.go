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
	n := nextInt()
	clouds := []int{}
	for i := 0; i < n; i++ {
		clouds = append(clouds, nextInt())
	}
	steps := 0
	i := 0
	for i < n-1 {
		if i < n-2 && clouds[i+2] == 0 {
			i += 2
		} else {
			i += 1
		}
		steps += 1
	}
	fmt.Println(steps)
}
