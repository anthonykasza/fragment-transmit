package main

import ("fmt")


func main() {

  // range'ing a string results in int32's, this is surprising
  s := "hello123"
  fmt.Println(fmt.Printf("%T", s))
  for _,ord := range s {
    fmt.Println(fmt.Printf("%T", ord), ord, string(ord))
  }

  fmt.Println("\n======\n")

  // range'ing an arrary of string results in strings, as expected
  a := []string{"hello", "123"}
  fmt.Println(fmt.Printf(a))
  for _,elem := range a {
    fmt.Println(fmt.Printf(elem), elem)
  }

  // it's funny to me that go cannot infer the type of the following string array
  //   a := []{"hello", "123"}
  // and that we must explicitly use
  //   a := []string{"hello", "123"}

}
