package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWordFrequencies(text string) map[string]int {
	frequencies := make(map[string]int)

	// Разбиваем строку на слова по пробельным символам
	rawWords := strings.Fields(text)

	for _, word := range rawWords {
		// 1. Приводим к нижнему регистру
		word = strings.ToLower(word)

		// 2. Очищаем слово от знаков препинания по краям (точки, запятые и т.д.)
		// Например, превращаем "привет," в "привет"
		cleanWord := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}
			return -1 // Удалить символ
		}, word)

		// Пропускаем пустые строки (если слово состояло только из знаков препинания)
		if cleanWord == "" {
			continue
		}

		// 3. Увеличиваем счетчик в мапе
		frequencies[cleanWord]++
	}

	return frequencies
}

func main() {
	text := "Привет мир! Привет Го. Как дела, мир? Мир прекрасен, привет!"

	fmt.Println("Анализируемый текст:", text)
	fmt.Println("--- Статистика ---")

	stats := countWordFrequencies(text)

	// Вывод результатов
	// Порядок вывода в мапах не гарантирован, но для демонстрации подойдет
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}
