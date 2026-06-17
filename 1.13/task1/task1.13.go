package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("не может быть равен 0")
	}
	return a / b, nil
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Тесты функции divide:")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("divide(10, 2) = %f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("divide(10, 0) вызвал ошибку: %v\n", err)
	} else {
		fmt.Printf("divide(10, 0) = %f\n", result)
	}

	result, err = divide(-10, 2)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("divide(-10, 2) = %f\n", result)
	}

	fmt.Println("\nТесты функции sumAll:")

	fmt.Println("sumAll(1, 2, 3) =", sumAll(1, 2, 3))           // 6
	fmt.Println("sumAll(10, -2, 4, 7) =", sumAll(10, -2, 4, 7)) // 19
	fmt.Println("sumAll(5) =", sumAll(5))                       // 5
	fmt.Println("sumAll() =", sumAll())                         // 0
	fmt.Println("sumAll(-1, -2, -3) =", sumAll(-1, -2, -3))     // -6
}
