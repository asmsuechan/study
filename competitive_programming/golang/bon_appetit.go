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

func calcActual(s []int, k int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i != k {
			sum += s[i]
		}
	}
	return sum / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()
	items := []int{}
	for i := 0; i < n; i++ {
		items = append(items, nextInt())
	}
	charged := nextInt()
	overcharged := charged - calcActual(items, k)
	if overcharged == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(overcharged)
	}
}
