package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mutex   sync.Mutex
}

// Deposit adds money to the account
func (a *Account) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance += amount
}

// Withdraw subtracts money from the account
func (a *Account) Withdraw(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance -= amount
}

func main() {
	account := Account{}

	var wg sync.WaitGroup
	wg.Add(4)

	// Deposit money concurrently
	go func() {
		defer wg.Done()
		account.Deposit(100)
	}()
	account.Deposit(100)

	// Withdraw money concurrently
	go func() {
		defer wg.Done()
		account.Withdraw(50)
	}()
	go func() {
		defer wg.Done()
		account.Deposit(200)
	}()

	// Withdraw money concurrently
	go func() {
		defer wg.Done()
		account.Withdraw(50)
	}()
	account.Withdraw(50)

	wg.Wait()

	fmt.Println("Final balance:", account.balance) // Final balance will be 50
}
