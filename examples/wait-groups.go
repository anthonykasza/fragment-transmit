package main
import (
  "fmt"
  "time"
  "sync"
)

// some work task
func worker(id int) {
  fmt.Println("worker start", id)
  time.Sleep(time.Second)
  fmt.Println("worker done", id)
}

func main() {
  // create a waitgroup
  var wg sync.WaitGroup

  for i := 1; i<=5; i++ {
    // i is not available inside the scope of the goroutine,
    //   so we set id to i and use it
    id := i

    // expand the number of things to wait on
    wg.Add(1)

    go func() {
      // shrink the number of things to wait on, but defer it until
      //   after some function is called
      defer wg.Done()

      // invoke some work inside the goroutine
      worker(id)
    }()
  }

  // wait for the waitgroup to finish, this happens when the number
  //   of calls to wg.Add() matches the number of calls to wg.Done()
  wg.Wait()
}
