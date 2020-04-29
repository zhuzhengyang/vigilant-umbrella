#!/bin/bash
echo "package gen

type Player struct {
	Name string
}" > ./src/gen/model.go
go build --buildmode=plugin -o sstcpserver1.so src/sstcpserver1/main.go

echo "package gen

type Player struct {
	Name string
	Age  int
}" > ./src/gen/model.go
go build --buildmode=plugin -o sstcpserver2.so src/sstcpserver2/main.go

echo "package gen

type Player struct {
	Name string
}" > ./src/gen/model.go

go run main.go