package main
import ("fmt")

func main() {
  recv_ch1 := make(chan int)
  recv_ch2 := make(chan string)

  go func() { recv_ch1 <- 1 }()
  go func() { recv_ch2 <- "two" }()

  // we want 6 attempts to read messages from our background tasks
  for i := 0; i < 6; i++ {
    select {
    case msg1 := <-recv_ch1:
      fmt.Println("received", msg1)
    case msg2 := <-recv_ch2:
      fmt.Println("received", msg2)
    // default says: if no messages are available, do thi instead
    default:
      fmt.Println("aint got no messages")
    }
  }
}
