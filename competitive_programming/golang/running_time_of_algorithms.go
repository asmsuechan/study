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

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println("")
}

func swap(a []int, fromIndex int, toIndex int) []int {
	prevValue := a[fromIndex]
	a[fromIndex] = a[toIndex]
	a[toIndex] = prevValue
	return a
}

func delete(slice []int, i int) (int, []int, error) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return 0, nil, fmt.Errorf("Error")
	}
	ans := make([]int, len(slice))
	copy(ans, slice)

	ans = append(slice[:i], slice[(i+1):]...)

	return ret, ans, nil
}

func move(a []int, index int, dest int) []int {
	d := 0
	d, a, _ = delete(a, index)
	t := []int{d}
	a = append(a[:dest], append(t, a[dest:]...)...)
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, nextInt())
	}
	numberOfShifts := 0
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			var j, dest int
			for j = i; j >= 0; j-- {
				if ar[i+1] < ar[j] {
					dest = j
					numberOfShifts += 1
				}
			}
			move(ar, i+1, dest)
		}
	}
	fmt.Println(numberOfShifts)
}
