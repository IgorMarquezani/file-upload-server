package main

import (
	"fmt"
	"net"
	"time"
)

func test() {
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    panic(err)
  }

  for _, addr := range addrs {
    fmt.Println(addr.String())
    fmt.Println(addr.Network())
  }

  time.Sleep(time.Second * 10)

  interfaces, err := net.Interfaces()
  if err != nil {
    panic(err)
  }

  for _, inter := range interfaces {
    fmt.Println(inter.Name)
    fmt.Println(inter.Addrs())
  }
}
