package main

import "fmt"

func main() {
  numbers := []int{ 1, 2 }

  printNumbers(numbers)

  result := reverseArray(numbers)

  printNumbers(result)
}

func printNumbers(numbers []int) {
  fmt.Println("numbers:")
  for i, el := range(numbers) {
    fmt.Println(fmt.Sprintf("%d: %d", i, el))
  }
}
