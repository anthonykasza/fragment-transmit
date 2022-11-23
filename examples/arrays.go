package main
import ("fmt"; "reflect")


// arrays are like slices but they don't have a capacity that can change

func main() {
  // an empty array of size 0
  a1 := [...]int{}
  fmt.Println(a1, reflect.TypeOf(a1))
}
