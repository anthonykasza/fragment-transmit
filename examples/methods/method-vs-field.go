package main
import ("fmt")

type FuncAsType func() string

type Receiver struct {
  // a member field as you'd expect
  member string

  // structs can also contain functions as fields
  MemberFunction FuncAsType
}

// struct methods are not the same as fields which are functions
func (r *Receiver) Method (arg_member string) {
  r.member = arg_member
  r.MemberFunction = func() string { return "altered value" }
}

func main() {
  my_receiver := Receiver{
    // pay attention to how members are defined. the following comments
    //   are incorrect:
    //   member := "hello123",
    //   member = "hello123",
    //   member: "hello123";
    member: "hello123",
    MemberFunction: func() string { return "default value" },
  }
  fmt.Println(my_receiver)
  my_receiver.MemberFunction()
  fmt.Println(my_receiver)
  fmt.Println(my_receiver.MemberFunction())

  my_receiver.Method("this argument becomes a member value")
  fmt.Println(my_receiver)
}
