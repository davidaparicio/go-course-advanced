#!/bin/sh

# failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH

#brew install graphviz
#go build .
#time cat BiblePass_part01.txt | ./anew -d BiblePass_part02.txt
go tool pprof cpu.pprof 
# Inside, try these commands
# top
# list main
# web

go tool pprof -http=":8000" cpu.pprof
# https://www.webperf.tips/tip/understanding-flamegraphs/