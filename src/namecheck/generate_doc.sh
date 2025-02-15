#!/usr/bin/env bash

#present main.slide
#open "http://127.0.0.1:3999"
go test -v ./twitter
#mkdir -p $HOME/go/{bin,src,pkg}
#go get golang.org/x/tools/cmd/godoc
#go install golang.org/x/tools/cmd/present && go install golang.org/x/tools/cmd/godoc
godoc -http=:6060 -play #-goroot=$HOME/go
open "http://127.0.0.1:6060"

# Why is GO111MODULE everywhere, and everything about Go Modules (updated with Go 1.17)
# https://maelvls.dev/go111module-everywhere/