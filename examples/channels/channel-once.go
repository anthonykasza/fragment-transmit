package main
import ("fmt")

// A Once is a channel of empty structs
type Once chan struct{}

func NewOnce() Once {
  o := make(Once, 1)
  // write an empty struct to the channel
  o <- struct{}{}
  return o
}

// assign a member function to the Once type
// the member function takes a func as an argument
//   which itself takes an empty struct as an arg
func (o Once) Do(f func(s struct{})) {
  // read from ourselves
  s, ok := <- o
  if !ok { return }
  // execute the function, passing it the empty struct we just read
  //   in from ourselves
  f(s)
  // close ourselves
  close(o)
}

func main() {
  my_o := NewOnce()
  my_o.Do(func(s struct{}) {
    fmt.Println("Doing it and doing it1", s)
  })

  // this never happens because we already closed ourselves at the end of Once
  my_o.Do(func(s struct{}) {
    fmt.Println("Doing it and doing it2", s)
  })
}
