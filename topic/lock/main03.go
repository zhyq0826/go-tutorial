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

func Balance2() int {
	mu.RLock()
	fmt.Println("Balance2 acquired lock")
	defer mu.RUnlock()
	time.Sleep(10 * time.Second)
	fmt.Println("Balance2 release lock")
	return balance
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
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
		fmt.Println("Balance2 == >", Balance2())
		wait <- 1
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Balance ==>", Balance())
	<-wait
}
