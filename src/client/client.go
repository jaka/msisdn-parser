package main

import "os"
import "net/rpc/jsonrpc"
import "msisdn"
import "fmt"

const SERVER_NAME string = ":12345"

type Results struct {
  Answer msisdn.Answer
}

func queryServer(msisdn string, results *Results) error {
  client, err := jsonrpc.Dial("tcp4", SERVER_NAME)
  if err == nil {
    err = client.Call("RPCMethods.ParseMSISDN", msisdn, results)
  }
  return err
}

func main() {

  var results Results

  if len(os.Args) != 2 {
    fmt.Println("Usage:", os.Args[0], "[msisdn number]")
    os.Exit(1)
  }
  var msisdn = os.Args[1]

  err := queryServer(msisdn, &results)

  if  err != nil {
    fmt.Println("Error:", err.Error())
  } else {
    answer := results.Answer
    fmt.Printf("Country identifier: %v\nCountry dialling code: %v\nMNO identifier: %v\nSubscriber number: %v\n",
      answer.CountryISO ,
      answer.CountryDial,
      answer.ProviderName,
      answer.SubscriberNumber)
  }
}
