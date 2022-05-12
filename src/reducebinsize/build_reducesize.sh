#!/bin/sh

go build -o main_init main.go 
go build -o main_sol1 -ldflags="-s -w" main.go 
cp main_sol1 main_sol2
#brew install upx
upx --brute main_sol2

ls -la main_init main_sol1 main_sol2