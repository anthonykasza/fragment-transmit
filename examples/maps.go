package main
import "fmt"

func main() {
  var nilmap map[int]string
  fmt.Println(nilmap)

  var emptymap = map[int]string{}
  fmt.Println(emptymap)

  var onemap = map[int]string {
   1: "hello",
  }
  fmt.Println(onemap)

  mademap := make(map[int]string)
  fmt.Println(mademap)
}
