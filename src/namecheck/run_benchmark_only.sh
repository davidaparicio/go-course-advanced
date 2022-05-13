#!/usr/bin/env bash

# This is somewhat strange in Golang. 
# To be honest it took me some time to figure a way out. 
# https://stackoverflow.com/a/49098923
# https://stackoverflow.com/a/48868785
# go help test
# RUN ALL TESTS
# go test ./...
# RUN BENCHMARK ONLY!
go test -v -run=^$ -bench . -benchmem -benchtime=3s ./twitter

# GO DOCUMENTATION
# https://go.dev/doc/code
# https://www.digitalocean.com/community/tutorials/how-to-write-your-first-program-in-go
# https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos
# https://documentation.help/Golang/how%20to%20write%20go%20code.html
