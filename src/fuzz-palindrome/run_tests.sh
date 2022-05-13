#!/usr/bin/env bash

echo "Let's Test"
go test -v
go test -coverprofile=coverage.out

echo "Let's Fuzz"
go test -fuzz FuzzIsPalindrome
#go test -fuzz FuzzIsPalindrome -fuzztime 15s
#go test -fuzz FuzzIsPalindrome_WithRune -fuzztime 15s 

echo "Let's Bench"
#go test -v -run=^$ -bench . -benchmem -benchtime=3s .