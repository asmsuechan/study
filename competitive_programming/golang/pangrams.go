package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func nextString() string {
	sc.Scan()
	n := sc.Text()
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

func convertAlphabetToNumber(s rune) int {
	switch s {
	case 'a':
		return 0
	case 'b':
		return 1
	case 'c':
		return 2
	case 'd':
		return 3
	case 'e':
		return 4
	case 'f':
		return 5
	case 'g':
		return 6
	case 'h':
		return 7
	case 'i':
		return 8
	case 'j':
		return 9
	case 'k':
		return 10
	case 'l':
		return 11
	case 'm':
		return 12
	case 'n':
		return 13
	case 'o':
		return 14
	case 'p':
		return 15
	case 'q':
		return 16
	case 'r':
		return 17
	case 's':
		return 18
	case 't':
		return 19
	case 'u':
		return 20
	case 'v':
		return 21
	case 'w':
		return 22
	case 'x':
		return 23
	case 'y':
		return 24
	case 'z':
		return 25
	case ' ':
		return 26
	default:
		return -1
	}
}

func main() {
	s := nextString()
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	n := make([]int, 26)
	for _, c := range s {
		n[convertAlphabetToNumber(c)]++
	}
	f := false
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			f = false
			break
		} else {
			f = true
		}
	}
	if f {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
