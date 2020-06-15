#!/bin/bash
# first
# go run main.go watch.go

projectpath=`pwd`
export GOPATH="${projectpath}:$GOPATH"
export GOBIN="${projectpath}/bin"
go build --buildmode=plugin -o ./plugin/sstcpserver1.so sstcpserver1
go build --buildmode=plugin -o ./plugin/sstcpserver2.so sstcpserver2
