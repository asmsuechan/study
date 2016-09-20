package main

import(
  "fmt"
  "math"
)

func main() {
  var (
    a string = "aaaaa"
    b = "bbbbb"
    c = "ccccc"
  )
  d := "ddddd"
  var i int
  pi := math.Pi
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(i)
  fmt.Println(pi)

  if a == "aaaaa" {
    fmt.Println("a is aaaaa")
  }

  for i = 0; i < 10; i++ {
    fmt.Println(i)
  }
}
