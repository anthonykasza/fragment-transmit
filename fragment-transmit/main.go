// A Dumb script that does data fragmentation and transmission

package main

import (
  "fmt"
  "net"
  "os"
  "math/rand"
)

const (
  PROTO = "tcp"
  FNAME = "./data-file"
  CHUNKS = 3
)

func constDests() [1]string {
  return [1]string{"127.0.0.1:9090"}
}

func main() {
  var DESTS = constDests()

  // shut up about unused imports
  var _ = net.Dial
  var _ = fmt.Println

  // determine data size and how to chunk it
  fi, err := os.Stat(FNAME)
  if err != nil { panic(err) }
  fsize:= fi.Size()
  chunk_size := fsize / CHUNKS;
  var sent_bytes int64 = 0;

  f, err := os.Open(FNAME)
  if err != nil { panic(err) }

  for sent_bytes < fsize {
    buf := make([]byte, chunk_size)
    read_size, err := f.Read(buf)
    if err != nil { panic(err) }

    // TODO - consider adding a chunk index field to the file data
    //        before transmitting so DESTS can reassemble

    // TODO - compress or encrypt the data before sending

    // select a drop location
    rand_idx := rand.Intn(len(DESTS))
    rand_dest := DESTS[rand_idx]

    // transfer the data
    conn, err := net.Dial(PROTO, rand_dest)
    if err != nil { panic(err) }
    _, err = conn.Write(buf[:read_size])
    conn.Close()

    // print for debugging
    fmt.Println("sent bytes:   ", string(buf[:read_size]))

    // adjust the file offset
    sent_bytes += chunk_size
  }
}
