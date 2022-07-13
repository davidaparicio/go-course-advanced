#!/bin/sh

go get -u honnef.co/go/tools/cmd/structlayout
go get -u honnef.co/go/tools/cmd/structlayout-pretty
go get -u honnef.co/go/tools/cmd/structlayout-optimize
go install honnef.co/go/tools/cmd/structlayout
go install honnef.co/go/tools/cmd/structlayout-pretty
go install honnef.co/go/tools/cmd/structlayout-optimize

#echo "BEFORE MEMORY OPTIMIZATION"
#structlayout -json ./main.go User | structlayout-pretty | tee old.mem
#echo "AFTER MEMORY OPTIMIZATION"
#structlayout -json ./main.go User2 | structlayout-pretty | tee new.mem

go get -u github.com/ajstarks/svgo/structlayout-svg
go install github.com/ajstarks/svgo/structlayout-svg

echo "BEFORE MEMORY OPTIMIZATION"
structlayout -json ./main.go User | structlayout-svg -t "User, before optimization" > user.svg
echo "AFTER MEMORY OPTIMIZATION"
structlayout -json ./main.go User2 | structlayout-svg -t "User, after optimization" > user2.svg

#echo "=============================================================================="
#echo ""
#echo "=========================="
#echo "=========================="
#echo "== MEMORY OPTIMIZATION =="
#echo "=========================="
#echo "=========================="
#https://difftastic.wilfred.me.uk/installation.html
#brew install difftastic
#difft old.mem new.mem

open user.svg
open user2.svg