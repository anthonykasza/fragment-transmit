package main

import (
  "fmt"
  "time"
)

func main() {
  // ticks are delivered repeatedly every time.Duration specified in
  //   time.NewTicker()
  // "C" is the channel on which the ticks are delivered.
  // type Ticker struct {
  //   C <-chan Time
  // }
  ticker := time.NewTicker(500 * time.Millisecond)

  // timers are the same as tickers but only happen once
  timer := time.NewTimer(100 * time.Millisecond)

  // done is a signaling channel used by main to signal
  //   to the goroutine that it's time to return
  done := make(chan bool)

  go func() {
    // this loop repeats forever until either:
    //   - the main process signals via the "done" channel
    //   - the Ticker delivers one of its periodic ticks
    for {
      select {
      case <-done:
        return
      case t := <-ticker.C:
        fmt.Println("Tick at", t)
      case t := <-timer.C:
        fmt.Println("Timer at", t)
      }
    }
  }()

  // we sleep the main process so that the ticker can tick
  //   and the timer can trigger
  time.Sleep(1600 * time.Millisecond)

  // we stop the ticker. this isn't strictly necessary and
  //   the following line can be removed from this example
  ticker.Stop()

  // we signal, from main, for the goroutine to return
  done <- true
}
