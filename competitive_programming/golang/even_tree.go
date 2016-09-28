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

func findChildren(a [][]int, n int) []int {
	children := []int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < 2; j++ {
			if a[i][j] == n {
				if j == 0 && a[i][1] > a[i][0] {
					children = append(children, a[i][1])
				} else if j == 1 && a[i][0] > a[i][1] {
					children = append(children, a[i][0])
				}
			}
		}
	}
	return children
}

func findParent(a [][]int, n int) int {
	var parent int
	for i := 0; i < len(a); i++ {
		for j := 0; j < 2; j++ {
			if a[i][j] == n {
				if j == 0 && a[i][1] < a[i][0] {
					parent = a[i][1]
				} else if j == 1 && a[i][0] < a[i][1] {
					parent = a[i][0]
				}
			}
		}
	}
	return parent
}

func isContained(a []int, n int) bool {
	f := false
	for _, v := range a {
		if v == n {
			f = true
		}
	}
	return f
}

func pop(a []int) ([]int, int) {
	v := a[len(a)-1]
	return a[:len(a)-1], v
}

func findVertices(a [][]int, n int) []int {
	children := []int{}
	unchecked := []int{}
	children = append(children, findChildren(a, n)...)
	unchecked = append(unchecked, findChildren(a, n)...)
	v := 0
	for len(unchecked) >= 0 {
		unchecked = append(unchecked, findChildren(a, v)...)
		if len(unchecked) == 0 {
			break
		}
		unchecked, v = pop(unchecked)
		children = append(children, findChildren(a, v)...)
	}
	return children
}

func findLeafs(a [][]int, numberOfAllVertices int) []int {
	leaf := []int{}
	for i := 1; i <= numberOfAllVertices; i++ {
		if len(findChildren(a, i)) == 0 {
			leaf = append(leaf, i)
		}
	}
	return leaf
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	tree := make([][]int, m)
	for i := 0; i < m; i++ {
		edge := []int{}
		edge = append(edge, nextInt())
		edge = append(edge, nextInt())
		//tree = append(tree[i], edge...)
		tree[i] = edge
	}
	numberOfForests := 0
	for i := 2; i < n; i++ {
		if len(findVertices(tree, i))%2 == 1 {
			numberOfForests++
		}
	}
	fmt.Println(numberOfForests)
}
