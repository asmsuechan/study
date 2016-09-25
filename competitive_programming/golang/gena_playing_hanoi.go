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

func push(a []int, n int) []int {
	return append(a, n)
}

func pop(a []int) ([]int, int) {
	length := len(a)
	popedVal := a[length-1]
	return a[:length-1], popedVal
}

func hanoi(a []int) {
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	discs := []int{}
	rod1, rod2, rod3, rod4 := []int{}, []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		discs = append(discs, nextInt())
	}
	//prepare 4 stacks for rod
	for i := n - 1; i >= 0; i-- {
		switch discs[i] {
		case 1:
			rod1 = push(rod1, i)
		case 2:
			rod2 = push(rod2, i)
		case 3:
			rod3 = push(rod3, i)
		case 4:
			rod4 = push(rod4, i)
		}
	}
	fmt.Println("rod1: ", rod1)
	fmt.Println("rod2: ", rod2)
	fmt.Println("rod3: ", rod3)
	fmt.Println("rod4: ", rod4)
}
