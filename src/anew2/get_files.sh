#!/bin/sh

wget https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/BiblePass/BiblePass_part01.txt
wget https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/BiblePass/BiblePass_part02.txt
 
#go install -v github.com/tomnomnom/anew@latest
time cat BiblePass_part01.txt | anew -d BiblePass_part02.txt