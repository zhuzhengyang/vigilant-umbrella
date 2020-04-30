package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
)

var PlayerNumber int

func AddPlayerNumber() {
	PlayerNumber++
}


func main() {
    go func() {
	    log.Println(http.ListenAndServe("localhost:6062", nil))
    }()
    ch := make(chan int, 0)
    AddPlayerNumber()
    fmt.Println("PlayerNumber1 ", PlayerNumber)

    select {
        case <- ch:
    }
}