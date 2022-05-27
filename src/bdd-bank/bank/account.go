package bank

type account struct {
	name    string
	balance int
}

func (a *account) deposit(amount int) {
	a.balance = a.balance + amount
}

func (a *account) withdraw(amount int) {
	a.balance = a.balance - amount
}

func NewAccount(name string) *account {
	return &account{name, 0}
}
