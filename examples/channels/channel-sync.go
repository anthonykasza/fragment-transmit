package main
import ("fmt")

func worker(unblocked chan bool) {
  fmt.Println("working...done")
  // unblock the channel by sending true
  //   we don't actually care if the worker sends back true or false
  //   just as long as we get a message across the channel
  // try changing the following true to false
  unblocked <- true
}

func main() {
  // we make a boolean channel with a buffer capacity of 1
  unblocked := make(chan bool, 1)

  go worker(unblocked)

  // block main() from exiting until we hear back from the worker
  <-unblocked
}