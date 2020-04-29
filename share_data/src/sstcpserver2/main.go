package main

import "fmt"

func init() {
	fmt.Println("init2")
}

func AddPlayerNumber(pn *int) {
	*pn ++
}
