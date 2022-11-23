package main
import (
  "fmt"
)

type IdxOrd struct {
  idx int
  ord rune
}

func encode(cleartext string) []rune {
  ciphertext := []rune{}
  recv_chan := make(chan IdxOrd, len(cleartext))

  for i,ord := range cleartext {
    ciphertext = append(ciphertext, 1)
    idx := i
    wi := ord
    fmt.Println("  encoding math:", string(ord), ord, (ord + 28) % 255)
    go func(idx int, wi rune, recv_chan chan IdxOrd) {
      foo := IdxOrd{idx, (wi + 28) % 255}
      recv_chan <- foo
    }(idx, wi, recv_chan)
  }

  for range cleartext {
    select {
    case encoded_idxord := <- recv_chan:
      ciphertext[encoded_idxord.idx] = encoded_idxord.ord
    }
  }

  return ciphertext
}

func main() {
  cleartext := "hello123"
  fmt.Println("cleartext:", cleartext)
  fmt.Println("encoded with goroutines:", encode(cleartext))
}
