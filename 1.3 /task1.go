package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Введите 5 чисел")

	var t1, t2, t3, t4, t5 float64
	fmt.Scan(&t1, &t2, &t3, &t4, &t5)
	t := []float64{t1, t2, t3, t4, t5} //запрашиваем ввод 5 чисел и записываем их в переменную t

	fmt.Println(t)

	// Сортировка массива в порядке убывания
	sort.Slice(t, func(i, j int) bool { return t[i] > t[j] })
	fmt.Println("Отсортированный массив по убыванию:", t)

	// Находим самое маленькое число
	min := t[0]
	fmt.Println("Самое большое число:", min)

	// Находим самое большое число
	max := t[4]
	fmt.Println("Самое маленькое число:", max)

	// Вычисляем среднее арифметическое
	var sum float64
	for _, value := range t {
		sum += value
	}

	sr := sum / float64(len(t))
	fmt.Printf("Среднее арифметическое: %.4f", sr)
}
