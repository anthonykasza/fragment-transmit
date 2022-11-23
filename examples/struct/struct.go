package main
import ("fmt")

type NestedStruct struct {
  nested_struct_integer1 int
  nested_struct_integer2 int
}

type Receiver struct {
  member string
}

func (r *Receiver) Method (arg_member string) {
  r.member = arg_member
}

func main() {
  // these are all incorrect ways to instantiate a new structure
  //
  // Convert a value to a Receiver type. However, the function expects
  //   an argument as input to convert
  //my_receiver := Receiver()
  //
  // Should the language be called "Go" or "Golang"? Search the interwebs
  //   for "empty braces go", then for "empty braces golang".
  // What are empty braces? It depends on where you see them:
  //    https://go.dev/ref/spec#Notation
  //my_receiver := {}
  //
  // I really thought this would convert an empty structure to a Receiver,
  //   but nahhhh
  //my_receiver := Receiver({})

  // correct, ayay
  my_receiver := Receiver{}
  fmt.Println(my_receiver)
}
