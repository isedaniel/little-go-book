package main

type Account struct {
	name    string
	balance int
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) SetBalance(amount int) {
	a.balance = amount
}

func NewAccount() *Account {
	return &Account{
		name:    "",
		balance: 0,
	}
}
