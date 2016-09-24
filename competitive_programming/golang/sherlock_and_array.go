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

func main() {
	sc.Split(bufio.ScanWords)
	T := nextInt()
	ans := make([]int, T)
	for i := 0; i < T; i++ {
		Ai := []int{}
		N := nextInt()
		sum := 0
		for j := 0; j < N; j++ {
			element := nextInt()
			Ai = append(Ai, element)
			sum += element
		}
		left := 0
		for k := 0; k < N; k++ {
			//left
			sum -= Ai[k]
			if left == sum {
				ans[i] = 1
				break
			}
			left += Ai[k]
		}
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
