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
	n := nextInt()
	var numberOfPositive, numberOfNegative, numberOfZeros int
	for i := 0; i < n; i++ {
		number := nextInt()
		switch {
		case number > 0:
			numberOfPositive += 1
		case number < 0:
			numberOfNegative += 1
		case number == 0:
			numberOfZeros += 1
		}
	}
	fmt.Printf("%1.6f\n", float64(numberOfPositive)/float64(n))
	fmt.Printf("%1.6f\n", float64(numberOfNegative)/float64(n))
	fmt.Printf("%1.6f\n", float64(numberOfZeros)/float64(n))
}
