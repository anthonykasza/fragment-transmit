package main
import ("fmt")

func main() {
  recv_ch1 := make(chan int)
  recv_ch2 := make(chan string)

  go func() { recv_ch1 <- 1 }()
  go func() { recv_ch2 <- "two" }()

  // we called 2 goroutines so we ned to wait for 
  //   2 channel messages to be recv'd
  //   if recv_ch1 was written to twice before recv_ch2
  //   has any writes, this loop would exit without
  //   reading any messages from recv_ch2
  for i := 0; i < 2; i++ {
    select {
    case msg1 := <-recv_ch1:
      fmt.Println("received", msg1)
    case msg2 := <-recv_ch2:
      fmt.Println("received", msg2)
    }
  }
}