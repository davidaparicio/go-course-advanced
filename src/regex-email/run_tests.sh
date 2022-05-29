#!/usr/bin/env bash

FuzzFUNC="Fuzz"

echo "Let's Test"
go test -v
go test -coverprofile=coverage.out

echo "Let's Fuzz"
go test -fuzz ${FuzzFUNC} -fuzztime 30s

echo "Let's Bench"
go test -v -run=^$ -bench . -benchmem -benchtime=3s .

# https://alicegg.tech/2019/03/09/gobdd.html
# https://gist.github.com/davidaparicio/9fa9e664b6eed47e054ee6de426abe1a
# https://github.com/cucumber/godog
# https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go
# https://medium.com/propertyfinder-engineering/golang-api-testing-with-godog-2de8944d2511
# https://www.refactoredtelegram.net/2021/06/parametrising-your-bdd-tests-in-go/

# https://tip.golang.org/doc/tutorial/fuzz