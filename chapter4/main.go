package main 

import "fmt"

func main() {
  fmt.Print("Enter the temperature in Fahrenheight: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := (input - 32) * 5 / 9
  fmt.Println(output, "degrees Celsius")
}