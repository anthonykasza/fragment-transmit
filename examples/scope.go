package main
import (
  "fmt"
)
const (
  FIVE int = 5
)

func main() {
  done := make(chan bool, FIVE*2)
  main_scoped_int := 4

  for for_scoped_int:=1; for_scoped_int<=FIVE; for_scoped_int++ {
    // for_scoped_int and main_scoped_int are inherited by reference
    fmt.Println(&main_scoped_int, &for_scoped_int, "for-loop's scope")

    go func() {
      // for_scoped_int and main_scoped_int are inherited by reference
      //   which means the for-loop will change for_scoped_int before
      //   this poor little gorouting can execute
      fmt.Println(&main_scoped_int, &for_scoped_int, "goroutine without arg")
      done <- true
    }()

    // function arguments are copied values
    go func(goroutine_scoped_int int) {
      // again, for_scoped_int's value is being changed from underneath
      //   this poor little goroutine. however, goroutine_scoped_int is a
      //   copy of for_scoped_int and the copy does not change
      fmt.Println(&main_scoped_int, &for_scoped_int, "goroutine with arg: ", &goroutine_scoped_int)
      done <- true
    }(for_scoped_int) // NOTICE THIS ARGUMENT
  }

  // wiat for all goroutines to signal done
  for n:=1; n<=FIVE*2; n++ { select { case <- done: } }
  // go let's us get away with some pretty goofy-looking syntax
}
