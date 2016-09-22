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
	t := nextInt()
	ans := []int{}
	for i := 0; i < t; i++ {
		nms := []int{}
		for j := 0; j < 3; j++ {
			nms = append(nms, nextInt())
		}
		id := nms[1] - 1 + nms[2]
		if id%nms[0] == 0 {
			ans = append(ans, nms[0])
		} else {
			ans = append(ans, id%nms[0])
		}
	}
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
