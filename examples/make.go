package main
import ("fmt")

func main() {
  // SLICE
  // type, length, and capacity
  slice_len := 5
  slice_cap := 6
  slice := make([]int, slice_len, slice_cap)
  fmt.Println("slice", slice)

  // MAP
  // why do we use a built-in function to make maps?
  //   mmap := map[int]int{}
  mmap := make(map[int]int)
  fmt.Println("map", mmap)

  // CHANNEL
  // this is pretty straightforward
  bufsize := 3
  c := make(chan int, bufsize)
  defer close(c)

  // arrays don't use make!
  //   array := [2]string{"one", "two"}
}
