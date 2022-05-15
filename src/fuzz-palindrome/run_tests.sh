#!/usr/bin/env bash

echo "Let's Test"
go test -v
go test -coverprofile=coverage.out

# echo "Let's Fuzz [1/4]"
# go test -fuzz FuzzIsPalindrome_1
# echo "-> FAILURE (it was expected!)"
# echo ""
# # ==> testdata/fuzz/FuzzIsPalindrome/15bd0 <==
# # go test fuzz v1
# # string("\xf5")

# echo "Let's print the bug"
# # tree testdata
# tail testdata/fuzz/FuzzIsPalindrome/* 
# # find . -type f -print -exec cat {} \;
# # cat testdata/fuzz/FuzzIsPalindrome/15bd0ea6a71e1138505c7d1011410810213bf67d2bdd5ac68d160fddb4d176c7
# echo ""
# echo "FIXED!"
# echo ""

# echo "Let's Fuzz [2/4]"
# go test -fuzz FuzzIsPalindrome_2 -fuzztime 15s
# echo "FIXED!"
# echo ""

# echo "Let's Fuzz [3/4]"
# go test -fuzz FuzzIsPalindrome_3 -fuzztime 15s 
# echo "FIXED!"
# echo ""

# echo "Let's Fuzz [4/4]"
# go test -fuzz FuzzIsPalindrome_4 -fuzztime 15s

echo "Let's Fuzz"
go test -fuzz FuzzIsPalindrome -fuzztime 15s

echo "Let's Bench"
go test -v -run=^$ -bench . -benchmem -benchtime=3s .