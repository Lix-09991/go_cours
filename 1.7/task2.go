package main

import (
	"fmt"
	"slices"
	"strings"
)

func AppendCity(slice *[]string, city string) {
	*slice = append(*slice, city)
}

func DeleteCity(slice *[]string, city string) {
	index := slices.IndexFunc(*slice, func(s string) bool {
		return strings.EqualFold(s, city)
	})
	if index != -1 {
		*slice = slices.Delete(*slice, index, index+1)
	}
}

func SearchCity(slice []string, city string) bool {
	return slices.IndexFunc(slice, func(s string) bool {
		return strings.EqualFold(s, city)
	}) != -1
}

func main() {
	mainslice := []string{
		"Владивосток",
		"Южно-Сахалинск",
		"Холмс",
		"Москва",
		"Сочи",
		"нью-Йорк", //краевой случай, должно быть true
	}

	fmt.Println("Исходный слайс:", mainslice)

	AppendCity(&mainslice, "Лондон")
	fmt.Println("Добавление нового города:", mainslice)

	DeleteCity(&mainslice, "Красноярск")
	fmt.Println("Удаление города по имени (ex. Красноярск):", mainslice)

	exists := SearchCity(mainslice, "Нью-Йорк")
	fmt.Println("Поиска города в списке:", exists)
}
