package main
import ("fmt")

func main() {
  data_queue := make(chan string, 2)
  signal := make(chan bool, 1)

  go func() {
    // range is used to iterate over the channel's messages until the sender closes the channel
    for elem := range data_queue {
      fmt.Println(elem)
    }
    // the for-loop terminates when main() closes the data_queue channel
    signal <- true
  }()

  // enque some data then close the channel
  // ONLY SENDERS SHALL CLOSE CHANNELS - SO LET IT BE WRITTEN
  data_queue <- "one"
  data_queue <- "two"
  close(data_queue)

  // have main() block until we receive a signal from our anonymous goroutine
  <-signal
}
