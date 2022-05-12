#!/bin/sh

#go install honnef.co/go/tools/cmd/structlayout
#go install honnef.co/go/tools/cmd/structlayout-pretty
#go install honnef.co/go/tools/cmd/structlayout-optimize
go get honnef.co/go/tools/cmd/structlayout
go get honnef.co/go/tools/cmd/structlayout-pretty
go get honnef.co/go/tools/cmd/structlayout-optimize
echo "BEFORE MEMORY OPTIMIZATION"
structlayout -json ./main.go User | structlayout-pretty | tee old.mem
echo "AFTER MEMORY OPTIMIZATION"
structlayout -json ./main.go User2 | structlayout-pretty | tee new.mem

echo "=============================================================================="
echo ""
echo "=========================="
echo "=========================="
echo "== MEMORY OPTIMIZATION =="
echo "=========================="
echo "=========================="
#https://difftastic.wilfred.me.uk/installation.html
#brew install difftastic
difft old.mem new.mem