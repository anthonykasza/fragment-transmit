package main
import ("fmt")

type Receiver struct {
  member string
}

func (r *Receiver) Method (arg_member string) {
  r.member = arg_member
}

func main() {
  my_receiver := Receiver{}
  fmt.Println(my_receiver)

  my_receiver.Method("this argument becomes a member value")
  fmt.Println(my_receiver)
}
