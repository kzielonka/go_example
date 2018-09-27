package main

func reverseArray(numbers []int) []int {
  newArray := make([]int, 0, len(numbers))
  for i := len(numbers) - 1; i >= 0; i-- {
    newArray = append(newArray, numbers[i])
  }
  return newArray
}
