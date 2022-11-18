package main
import ("fmt"; "reflect")


// arrays are like slices but they don't have a capacity that can change

func main() {
  // an empty array of size 0
  s1 := [...]int{}
  fmt.Println(s1, reflect.TypeOf(s1))
}
