package main

import "fmt"

func init() {
	fmt.Println("init1")
}

func AddPlayerNumber(pn *int) {
	*pn ++
}
