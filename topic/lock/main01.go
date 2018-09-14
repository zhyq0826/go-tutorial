package main

import (
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func main() {
	balance = 100
	Withdraw(110)
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func deposit(amount int) {
	balance = balance + amount
}
