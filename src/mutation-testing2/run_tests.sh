#!/usr/bin/env bash

#FuzzFUNC="Fuzz" #"FuzzReverse"

echo "Let's Test"
go test -v
go test -timeout 30s -coverprofile=coverage.out

echo "Let's Mutate (Beware Zombies!)"
# The goal is : The mutation score is 1.000000
# (3 passed, 0 failed, 0 duplicated, 0 skipped, total is 3)

#go get github.com/zimmski/go-mutesting/... && go install github.com/zimmski/go-mutesting/...
go-mutesting .
# https://dev.to/guilhermeguitte/mutation-testing-on-go-1lbf

#echo "Let's Fuzz"
#go test -fuzz ${FuzzFUNC} -fuzztime 15s

#echo "Let's Bench"
#go test -v -run=^$ -bench . -benchmem -benchtime=3s .