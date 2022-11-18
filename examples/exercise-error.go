package main

import (
  "fmt"
  "math"
  "errors"
)

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New(fmt.Sprintf("cannot Sqrt negative number: %v", float64(x)))
  }
  return math.Sqrt(x), nil
}

func main() {
  // we could cach the error with: i, err = Sqrt(-2)
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
