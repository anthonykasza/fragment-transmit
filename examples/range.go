package main

import ("fmt"; "reflect")


func main() {

  // range'ing a string results in int32's, this is surprising
  s := "hello123"
  fmt.Println(reflect.TypeOf(s))
  for _,ord := range s {
    fmt.Println(reflect.TypeOf(ord), ord, string(ord))
  }

  fmt.Println("\n======\n")

  // range'ing an arrary of string results in strings, as expected
  a := []string{"hello", "123"}
  fmt.Println(a)
  for _,elem := range a {
    fmt.Println(elem)
  }

  // it's funny to me that go cannot infer the type of the following string array
  //   a := []{"hello", "123"}
  // and that we must explicitly use
  //   a := []string{"hello", "123"}

}
