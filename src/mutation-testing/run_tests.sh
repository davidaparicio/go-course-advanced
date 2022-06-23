#!/usr/bin/env bash

#FuzzFUNC="Fuzz" #"FuzzReverse"

echo "Let's Test"
go test -v
go test -timeout 30s -coverprofile=coverage.out
#ONE LINER go test -timeout 30s -coverprofile=coverage.out . #github.com/davidaparicio/go-course-advanced/src/mutation-testing

echo "Let's Mutate"
#go get github.com/zimmski/go-mutesting/... && go install github.com/zimmski/go-mutesting/...
go-mutesting .

#echo "Let's Fuzz"
#go test -fuzz ${FuzzFUNC} -fuzztime 15s

#echo "Let's Bench"
#go test -v -run=^$ -bench . -benchmem -benchtime=3s .