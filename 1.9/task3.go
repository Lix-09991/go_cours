package main

import (
	"fmt"
)

// Структура BankAccount
type BankAccount struct {
	Owner   string
	Balance float64
}

// Метод Deposit
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

// Метод Withdraw
func (b *BankAccount) Withdraw(amount float64) bool {
	if b.Balance >= amount {
		b.Balance -= amount
		return true
	}
	fmt.Println("Недостаточно средств!")
	return false
}

func main() {
	// Создаем счет
	account := BankAccount{
		Owner:   "Анна Петрова",
		Balance: 1000.0,
	}

	// Пополняем счет
	account.Deposit(500.0)
	fmt.Printf("Баланс после пополнения: %.2f\n", account.Balance)

	// Пытаемся снять деньги
	if account.Withdraw(300.0) {
		fmt.Printf("Баланс после снятия: %.2f\n", account.Balance)
	}
}
