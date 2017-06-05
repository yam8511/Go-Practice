package main

import "fmt"

func main()  {
  for line := 10; line >= 2; line-- {
    for num := 1; num <= 10; num++ {
      if line >= num {
        fmt.Print("*")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println("")
  }
}
