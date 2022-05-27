package bank

import (
	"fmt"
	"os"
)

const (
	//https://go.dev/doc/effective_go#mixed-caps
	exitOk        = iota //0
	exitWithError        //1
)

func main() {
	defer fmt.Println("NoLog")
	acc := NewAccount("Alice")
	acc.deposit(100)
	acc.withdraw(50)
	fmt.Printf("Welcome %s to a bank for a changing world, you have %d\n", acc.name, acc.balance)
	os.Exit(exitOk)
}
