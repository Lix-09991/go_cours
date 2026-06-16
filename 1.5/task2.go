package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountVowels(s string) int {
	count := 0
	vowels := "аеёиоуыэюя"
	//приведем строку к одному виду, так как заглавные != строчные
	s = strings.ToLower(s)
	//цикл по каждой руне value строки x
	for _, value := range s {
		if strings.ContainsRune(vowels, value) {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	result := CountVowels(input)
	fmt.Printf("В строке \"%s\" %d гласных букв.\n", input, result)
}
