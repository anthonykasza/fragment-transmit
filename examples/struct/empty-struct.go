package main
import "fmt"

type EmptyStruct struct{}

func main() {
  var my_es EmptyStruct
  fmt.Println(my_es)
  fmt.Println(struct{}{})

  type MainScopedEmptyStruct struct{}
  var my_mses MainScopedEmptyStruct
  fmt.Println(my_mses)

  var a,b struct{}
  fmt.Println(a, b, a == b)
  fmt.Println(&a, &b, &a == &b)

  // This fails because struct{}{}, EmptyStruct, and MainScopedEmptyStruct
  //   are all diferrent types even though they are essentially fungible
  //fmt.Println(&a == &b, &a == &my_es, &a == &my_mses)
}
