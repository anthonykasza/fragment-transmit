package main
import ("fmt")

// the first stage of some data pipeline
func stage1(send_ch chan<- string, msg string) {
  send_ch <-msg + "Stage1"
}

// the second stage of some data pipeline
func stage2(recv_ch <-chan string, send_ch chan<- string) {
  // what a gooey syntax, https://stackoverflow.com/a/60030832
  send_ch<- <-recv_ch + "Stage2"
}

// creating the data pipeline and pushing data through it
func main() {
  send_ch := make(chan string, 1)
  recv_ch := make(chan string, 1)

  go stage1(send_ch, "Main")
  go stage2(send_ch, recv_ch)
  msg := <-recv_ch
  fmt.Println(msg)
}