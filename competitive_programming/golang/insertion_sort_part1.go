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

func pop(slice []int) (int, []int) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println("")
}

func main() {
	sc.Split(bufio.ScanWords)
	size := nextInt()
	ar := []int{}
	for i := 0; i < size; i++ {
		ar = append(ar, nextInt())
	}
	v := ar[size-1]
	f := 1
	ar[size-1] = ar[size-2]
	for i := len(ar) - 2; i >= 0; i-- {
		if ar[i] > v {
			ar[i+1] = ar[i]
			printArray(ar)
		} else {
			f = 0
			ar[i+1] = v
			printArray(ar)
			break
		}
	}
	if f == 1 {
		ar[0] = v
		printArray(ar)
	}
}
