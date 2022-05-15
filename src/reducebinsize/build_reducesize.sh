#!/bin/sh

go build -o exec_init main.go 
go build -o exec_sol1 -ldflags="-s -w" main.go 
cp exec_sol1 exec_sol2
#brew install upx
upx --brute exec_sol2

ls -la exec_init exec_sol1 exec_sol2