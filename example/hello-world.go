package main

import "fmt"
import "github.com/scottserok/qrencode"

func main() {
  qrencode.Init()
  filepath, err := qrencode.EncodeString("hello, world")
  if err != nil {
    panic("Error making QR Code")
  } else {
    fmt.Println(filepath)
    qrencode.Destroy(filepath)
  }
}
