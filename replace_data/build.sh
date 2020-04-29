#!/bin/bash
projectpath=`pwd`
export GOPATH="${projectpath}:$GOPATH"
export GOBIN="${projectpath}/bin"
go build --buildmode=plugin -o sstcpserver1.so sstcpserver1
go build --buildmode=plugin -o sstcpserver2.so sstcpserver2
go run main.go