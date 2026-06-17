package main

import (
	"testing"
)

func TestDivide(t *testing.T) {

	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Ожидалось nil, но получена ошибка: %v", err)
	}
	if result != 5 {
		t.Errorf("Ожидалось 5, но получено %f", result)
	}

	// деление с отрицательными числами
	result, err = divide(-10, 2)
	if err != nil {
		t.Errorf("Ожидалось nil, но получена ошибка: %v", err)
	}
	if result != -5 {
		t.Errorf("Ожидалось -5, но получено %f", result)
	}

	//деление на ноль (должна вернуть ошибку)
	result, err = divide(10, 0)
	if err == nil {
		t.Errorf("Ожидалась ошибка при делении на ноль, но её нет")
	}
	if err.Error() != "division by zero" {
		t.Errorf("Ожидалась ошибка 'division by zero', но получена: %v", err)
	}
	if result != 0 {
		t.Errorf("Ожидалось 0 при делении на ноль, но получено %f", result)
	}

	//  деление с дробными числами
	result, err = divide(7.5, 2.5)
	if err != nil {
		t.Errorf("Ожидалось nil, но получена ошибка: %v", err)
	}
	if result != 3 {
		t.Errorf("Ожидалось 3, но получено %f", result)
	}

	//ноль на число (должно работать)
	result, err = divide(0, 5)
	if err != nil {
		t.Errorf("Ожидалось nil, но получена ошибка: %v", err)
	}
	if result != 0 {
		t.Errorf("Ожидалось 0, но получено %f", result)
	}
}
