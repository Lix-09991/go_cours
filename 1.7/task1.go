/*package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func FillArray(arr *[10]int) {

	t := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range arr {
		arr[index] = t.Intn(100) + 1
	} //заполняем переданный массив случайными числами
}

func main() {
	var mass [10]int
	FillArray(&mass)
	fmt.Println(mass)

	slice := make([]int, len(mass))
	copy(slice, mass[:]) //копируем массив в слайс, чтобы не изменять исходный

	sort.Ints(slice)

	fmt.Println("Начальный массив: ", mass)
	fmt.Println("Отсортированный массив: ", slice)
}
