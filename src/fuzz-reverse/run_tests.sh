#!/usr/bin/env bash

echo "Let's Test"
go test -v
go test -coverprofile=coverage.out

echo "Let's Fuzz"
go test -fuzz FuzzFUNC -fuzztime 15s

echo "Let's Bench"
go test -v -run=^$ -bench . -benchmem -benchtime=3s .

# https://tip.golang.org/doc/fuzz/
# https://tip.golang.org/doc/tutorial/fuzz
# https://go.dev/blog/subtests
# https://julien.ponge.org/blog/playing-with-test-fuzzing-in-go/