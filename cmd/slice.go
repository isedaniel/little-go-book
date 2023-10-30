package main

import "fmt"

func arrayCap() {
	accounts := []Account{
		{"Dani", 100},
		{"Gabo", 200},
		{"Thompson", 300},
	}
	balances := extractBalances(accounts)
	fmt.Println(balances)
}

func extractBalances(a []Account) []int {
	balances := []int{}
	for _, acc := range a {
		balances = append(balances, acc.balance)
	}
	return balances
}
