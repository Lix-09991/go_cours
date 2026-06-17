package main

import (
	"testing"
)

func TestSumAll(t *testing.T) {
	// Тест 1: простая сумма (1, 2, 3) = 6
	result := sumAll(1, 2, 3)
	if result != 6 {
		t.Errorf("Ожидалось 6, но получено %d", result)
	}

	// Тест 2: сумма с отрицательными числами (10, -2, 4, 7) = 19
	result = sumAll(10, -2, 4, 7)
	if result != 19 {
		t.Errorf("Ожидалось 19, но получено %d", result)
	}

	// Тест 3: одиночное число
	result = sumAll(5)
	if result != 5 {
		t.Errorf("Ожидалось 5, но получено %d", result)
	}

	// Тест 4: пустой список (0 чисел) = 0
	result = sumAll()
	if result != 0 {
		t.Errorf("Ожидалось 0 для пустого списка, но получено %d", result)
	}

	// Тест 5: все отрицательные числа
	result = sumAll(-1, -2, -3)
	if result != -6 {
		t.Errorf("Ожидалось -6, но получено %d", result)
	}

	// Тест 6: смесь с нулями
	result = sumAll(0, 5, 0, 10)
	if result != 15 {
		t.Errorf("Ожидалось 15, но получено %d", result)
	}

	// Тест 7: большие числа
	result = sumAll(1000000, 2000000, 3000000)
	if result != 6000000 {
		t.Errorf("Ожидалось 6000000, но получено %d", result)
	}
}
