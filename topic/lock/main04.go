package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.RWMutex
	balance int
)

func Balance() int {
	fmt.Println("Balance wait for another goroutine release lock")
	mu.RLock()
	fmt.Println("Balance acquired lock")
	defer mu.RUnlock()
	return balance
}

func Deposit(amount int) {
	mu.Lock()
	fmt.Println("Deposit acquired lock")
	defer mu.Unlock()
	time.Sleep(5 * time.Second)
	deposit(amount)
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

func main() {
	balance = 100
	wait := make(chan int)
	go func() {
		Deposit(100)
		wait <- 1
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Balance ==>", Balance())
	<-wait
}
