package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	//Для корректного подсчёта символов, особенно с кириллицей используем utf8.RuneCountInString.
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	text, _ := reader.ReadString('\n') //  ReadString принимает строку и ошибку, ошибку не обрабатываем строку записываем в переменную переменную text

	// Считаем количество символов (рун), а не байтов
	count := utf8.RuneCountInString(text)
	fmt.Printf("Количество символов: %d\n", count)
}
