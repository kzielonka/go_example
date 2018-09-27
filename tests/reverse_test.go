package main

import "fmt"
import "testing"

func TestReverseingEmptyArray(t *testing.T) {
  numbers := []int{}

  result := reverseArray(numbers)

  if (len(result) != 0) {
    t.Error("expected to return empty array")
  }
}

func TestReversingOneElementArray(t *testing.T) {
  numbers := []int{23}

  result := reverseArray(numbers)

  if (len(result) != 1) {
    t.Error(fmt.Sprintf("expected to return one element array but is %d", len(result)))
  }

  if (result[0] != 23) {
    t.Error(fmt.Sprintf("expected to return same element but is %d", result[0]))
  }
}

func TestReversingTwoElementsArray(t *testing.T) {
  numbers := []int{23, 58}

  result := reverseArray(numbers)

  if (len(result) != 2) {
    t.Error(fmt.Sprintf("expected to return two elements array but is %d", len(result)))
  }

  if (result[0] != 58) {
    t.Error(fmt.Sprintf("expected to valid first element but is %d", result[0]))
  }

  if (result[1] != 23) {
    t.Error(fmt.Sprintf("expected to valid second element but is %d", result[1]))
  }
}
