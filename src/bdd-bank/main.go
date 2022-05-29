package main

import (
	"fmt"

	"github.com/davidaparicio/go-course-advanced/src/bdd-godog/bank"
)

const (
	//https://go.dev/doc/effective_go#mixed-caps
	exitOk        = iota //0
	exitWithError        //1
)

func main() {
	//defer fmt.Println("NoLog")
	acc := bank.NewAccount("Alice")
	acc.Deposit(100)
	acc.Withdraw(50)
	fmt.Printf("Welcome %s to a bank for a changing world, you have %d$\n", acc.Name(), acc.Balance())
	//os.Exit(exitOk)
}
