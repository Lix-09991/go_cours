package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkParentheses(s string) (bool, int, int) {
	openCount, closeCount := 0, 0
	balance := 0 

	for _, ch := range s {
		switch ch {
		case '(':
			openCount++
			balance++
		case ')':
			closeCount++
			balance--
			
			if balance < 0 {
				return false, openCount, closeCount
			}
		}
	}

	return balance == 0, openCount, closeCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку-формулу: ")
	formula, _ := reader.ReadString('\n')

	if len(formula) > 0 && formula[len(formula)-1] == '\n' {
		formula = formula[:len(formula)-1]
	}

	isValid, open, close := checkParentheses(formula)

	if isValid {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", open, close)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающие\n", open, close)
	}
}
