package main

import "os"
import "fmt"
import "net"
import "net/rpc"
import "net/rpc/jsonrpc"
import "msisdn"
import "errors"

type RPCMethods struct {}
type Results struct {
  Answer msisdn.Answer
}

func (m *RPCMethods) ParseMSISDN (v string, results *Results) error {
  var ok bool
  (*results).Answer, ok = msisdn.ParseMSISDN(v)
  if ok {
    return nil
  }
  return errors.New("Invalid MSISDN")
}

func checkError(err error) {
  if err != nil {
    fmt.Println("Fatal error ", err.Error())
    os.Exit(1)
  }
}

func main() {

  rpc.Register(new(RPCMethods))

  tcpAddr, err := net.ResolveTCPAddr("tcp4", ":12345")
  checkError(err)

  listener, err := net.ListenTCP("tcp4", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    jsonrpc.ServeConn(conn)
  }

}
