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

func sum(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	scores := scanArray(n)
	numberOfCandiesToEachChild := make([]int, n)
	numberOfCandiesToEachChild[0] = 1
	for i := 1; i < n; i++ {
		numberOfCandiesToEachChild[i] = 1
		if scores[i-1] < scores[i] {
			numberOfCandiesToEachChild[i] = numberOfCandiesToEachChild[i-1] + 1
		} else {
			numberOfCandiesToEachChild[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		//numberOfCandiesToEachChild[i] = 1
		if scores[i] > scores[i+1] && numberOfCandiesToEachChild[i] <= numberOfCandiesToEachChild[i+1] {
			numberOfCandiesToEachChild[i] = numberOfCandiesToEachChild[i+1] + 1
		} else {
			//numberOfCandiesToEachChild[i] = 1
		}
	}
	//fmt.Println(numberOfCandiesToEachChild)
	fmt.Println(sum(numberOfCandiesToEachChild))
}
