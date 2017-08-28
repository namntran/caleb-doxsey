package main 

import "fmt"

func main() {
  fmt.Print("Enter a number between 1 to 20: ")
  var input float64
  fmt.Scanf("%f", &input)

  i := input 

  if i > 10 {
    fmt.Println("Yuge")
  } else {
    fmt.Println ("Tiny")
  }
}