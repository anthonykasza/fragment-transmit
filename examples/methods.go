package main
import (
  "fmt"
)

type MyThing struct { S string; F float64 }

// view the receiver's value but don't modify it
func (mt MyThing) GetF () float64 {
  return mt.F
}

// modify the receiver's members by refering to them
func (mt *MyThing) SetF (f float64) {
  mt.F = f
}

// make a copy of the receiver and modify the copy's value
func (mt MyThing) SetFWrong (f float64) {
  mt.F = f
}

func main() {
  myThing := MyThing{"myString", 0}
  fmt.Println(myThing)
  fmt.Println(myThing.F)
  fmt.Println(myThing.GetF())

  fmt.Println("\n========\n")

  myThing.SetFWrong(999)
  fmt.Println(myThing)

  fmt.Println("\n========\n")

  myThing.SetF(10)
  fmt.Println(myThing)
}
