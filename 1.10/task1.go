package main

import (
	"fmt"
	"math"
)

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	CardHolder string
	CardNumber string
}

func (c CreditCard) Process(amount float64) string {
	return fmt.Sprintf("CreditCard: платеж %.2f успешно обработан для %s (****%s)",
		amount, c.CardHolder, last4(c.CardNumber))
}

type CryptoWallet struct {
	WalletAddress string
	Currency      string
}

func (c CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("CryptoWallet: платеж %.2f %s успешно отправлен на адрес %s",
		amount, c.Currency, c.WalletAddress)
}

func last4(s string) string {
	if len(s) <= 4 {
		return s
	}
	return s[len(s)-4:]
}

func main() {
	processors := []PaymentProcessor{
		CreditCard{
			CardHolder: "Ivan Petrov",
			CardNumber: "1234567812345678",
		},
		CryptoWallet{
			WalletAddress: "0xA1B2C3D4E5F6",
			Currency:      "USDT",
		},
	}

	for _, p := range processors {
		fmt.Println(p.Process(150.75))
	}

	fmt.Println("Пример числа:", math.Round(10.4))
}
