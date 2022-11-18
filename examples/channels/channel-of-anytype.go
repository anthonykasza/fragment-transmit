package main
import ("fmt")

// send various types across a single, unidirectional, channel
func send_anything(send_ch chan<- interface{}, i interface{}) {
  send_ch <- i
}

func main() {
  recv_ch := make(chan interface{})

  go send_anything(recv_ch, 1)
  go send_anything(recv_ch, "two")
  go send_anything(recv_ch, fmt.Println)
  go send_anything(recv_ch, recv_ch)

  // wait to receive 4 "things" from the channel
  for i := 0; i < 4; i++ {
    fmt.Println(<-recv_ch)
  }
}
