#!/bin/bash 

go mod tidy 
sync 
go get -v .
sync 
env GOOS=linux GOARCH=amd64 CGO_Enabled=0 go build -v .