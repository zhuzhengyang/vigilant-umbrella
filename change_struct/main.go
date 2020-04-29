package main

import (
    "changestruct/src/gen"
    "fmt"
    "os"
    "plugin"
)


func main() {
    player := &gen.Player{}
    p, err := plugin.Open("./sstcpserver1.so")
    if err != nil {
        fmt.Println("error open plugin1: ", err)
        os.Exit(-1)
    }
    s, err := p.Lookup("AddPlayerNumber")
    if err != nil {
       fmt.Println("error lookup AddPlayerNumber1: ", err)
       os.Exit(-1)
    }
    if AddPlayerNumber, ok := s.(func(*gen.Player)); ok {
       AddPlayerNumber(player)
    }
    fmt.Println(player)

    p, err = plugin.Open("./sstcpserver2.so")
    if err != nil {
       fmt.Println("error open plugin2: ", err)
       os.Exit(-1)
    }
    s, err = p.Lookup("AddPlayerNumber")
    if err != nil {
       fmt.Println("error lookup PlayerNumber2: ", err)
       os.Exit(-1)
    }
    if AddPlayerNumber, ok := s.(func(*gen.Player)); ok {
       AddPlayerNumber(player)
    }
    fmt.Println(player)
}