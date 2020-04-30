package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "plugin"
    _ "net/http/pprof"
)

func main() {
    go func() {
	    log.Println(http.ListenAndServe("localhost:6061", nil))
    }()
    ch := make(chan int, 0)
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
    if AddPlayerNumber, ok := s.(func()); ok {
       AddPlayerNumber()
    }
    s, err = p.Lookup("PlayerNumber")
    if err != nil {
       fmt.Println("error lookup PlayerNumber1: ", err)
       os.Exit(-1)
    }
    if PlayerNumber, ok := s.(*int); ok {
       fmt.Println("PlayerNumber1 ", *PlayerNumber)
    } else {
    	fmt.Println("PlayerNumber1 error")
    }

    p, err = plugin.Open("./sstcpserver2.so")
    if err != nil {
       fmt.Println("error open plugin2: ", err)
       os.Exit(-1)
    }
    s, err = p.Lookup("PlayerNumber")
    if err != nil {
       fmt.Println("error lookup PlayerNumber2: ", err)
       os.Exit(-1)
    }
    if PlayerNumber, ok := s.(*int); ok {
       fmt.Println("PlayerNumber2 ", *PlayerNumber)
    }

    select {
        case <- ch:
    }
}