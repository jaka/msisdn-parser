package main

import "os"
import "net/rpc/jsonrpc"
import "msisdn"
import "fmt"

const SERVER_NAME string = ":12345"

type Results struct {
  Answer msisdn.Answer
}

func queryServer(server_name, msisdn string, results *Results) error {
  client, err := jsonrpc.Dial("tcp4", server_name)
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

  server_name := os.Getenv("SERVER")
  if server_name == "" {
    server_name = SERVER_NAME
  }
  err := queryServer(server_name, msisdn, &results)

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
