#!/usr/bin/env bash

go test -v ./twitter
#go get golang.org/x/tools/cmd/godoc
godoc -http=:6060 -play #-goroot=$HOME/go
open "http://127.0.0.1:6060"
