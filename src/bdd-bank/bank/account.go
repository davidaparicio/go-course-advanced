package bank

type account struct {
	name    string
	balance int
}

func (a *account) Deposit(amount int) {
	a.balance = a.balance + amount
}

func (a *account) Withdraw(amount int) {
	a.balance = a.balance - amount
}

func (a *account) Name() string {
	return a.name
}

func (a *account) Balance() int {
	return a.balance
}
func NewAccount(name string) *account {
	return &account{name, 0}
}
