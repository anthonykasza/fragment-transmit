package main
import (
  "fmt"
  "time"
)

func main() {
  recv_ch1 := make(chan string, 1)  
  go func() {
    time.Sleep(2 * time.Second)
    recv_ch1 <- "result 1"
  }()
  select {
  case msg := <-recv_ch1:
    fmt.Println(msg)
  // time.After is function that returns a channel write of type Time
  case foo := <-time.After(1 * time.Second):
    fmt.Println("timeout 1", foo)
  }


  recv_ch2 := make(chan string, 1)
  go func() {
    time.Sleep(2 * time.Second)
    recv_ch2 <- "result 2"
  }()
  select {
  case msg := <-recv_ch2:
    fmt.Println(msg)
  // this case omits the data written to the channel as it's a signal channel and we don't care what the value is
  //   compare to line 17
  case <-time.After(3 * time.Second):
    fmt.Println("timeout 2")
  }
}
