package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanBigInt() int64{
  sc.Scan()
  n, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return int64(n)
}

func scanInt() int{
  sc.Scan()
  n, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return n
}

func main() {
  sc.Split(bufio.ScanWords)
  n := scanInt()
  bigSum := int64(0)
  for i := 0; i < n; i++ {
    bigSum += scanBigInt()
  }
  fmt.Println(bigSum)
}
