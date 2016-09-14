package main

import "fmt"
import "msisdn"

func main() {

  answer, ok := msisdn.ParseMSISDN("+38631700700")

  if ok {
    fmt.Println(answer)
  }
}
