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

var numberOfCells int = 0

func floodFill(a [][]int, i int, j int) {
	if i >= 0 && j >= 0 && i < len(a) && j < len(a[i]) {
		if a[i][j] == 1 {
			a[i][j] = 2
			numberOfCells++
			floodFill(a, i+1, j)
			floodFill(a, i-1, j)
			floodFill(a, i, j+1)
			floodFill(a, i, j-1)
			floodFill(a, i+1, j+1)
			floodFill(a, i+1, j-1)
			floodFill(a, i-1, j-1)
			floodFill(a, i-1, j+1)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = append(matrix[i], scanArray(m)...)
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			numberOfCells = 0
			floodFill(matrix, i, j)
			if max < numberOfCells {
				max = numberOfCells
			}
		}
	}
	fmt.Println(max)
}
