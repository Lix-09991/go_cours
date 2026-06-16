package main

import (
	"fmt"
)

// Структура Book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Метод GetInfo
func (b Book) GetInfo() string {
	return fmt.Sprintf("%q by %s, %d", b.Title, b.Author, b.Year)
}

func main() {
	// Создаем экземпляр книги
	book := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	// Выводим информацию
	fmt.Println(book.GetInfo())
}
