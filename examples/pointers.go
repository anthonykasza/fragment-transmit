package main
import "fmt"

func foo(t *int) {
  k := 101
  // t is a copy, change the copy all you want
  t = &k
}

func bar(t *int) {
  k := 102
  // the ptr can only point at new memory
  *t = k
}

func baz(t int) {
  k := 103
  // t's value is scoped to baz(), again because its a copy
  t = k
}

func main() {
  i := 100
  fmt.Println(i)

  foo(&i)
  fmt.Println(i)

  bar(&i)
  fmt.Println(i)

  baz(i)
  fmt.Println(i)
}


// pointer types can point to the address of any value
// pointer types have a fixed size
//   32-bit machine: 32bits (4 bytes)
//   64-bit machine: 64bits (8 bytes)
// pointer default values are nil
// &value results in the memory address for the variable named "value"
