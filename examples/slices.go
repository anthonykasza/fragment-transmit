package main
import ("fmt"; "reflect")

// slices are like arrays but slices' capacity can change,
//   allowing them to grow and shrink dynamically

func main() {
  slc_length := 10
  slc_capacity := 12
  slice := make([]int, slc_length, slc_capacity)

  fmt.Println(reflect.TypeOf(slice))
  fmt.Println(len(slice), cap(slice), slice)
}
