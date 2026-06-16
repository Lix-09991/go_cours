package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func capitalizeWord(word string) string {
	if word == "" {
		return ""
	}

	// ищем первую руну и оставшуюся часть
	r, size := utf8.DecodeRuneInString(word)
	if r == utf8.RuneError && size == 1 {

		return word
	}

	first := string(unicode.ToUpper(r))
	rest := word[size:]
	// А оставшуюся часть к нижнему регистру
	rest = strings.Map(func(r rune) rune { return unicode.ToLower(r) }, rest)

	return first + rest
}

func capitalizeWords(s string) string {
	if s == "" {
		return s
	}

	parts := strings.Fields(s)
	for i, w := range parts {
		parts[i] = capitalizeWord(w)
	}
	return strings.Join(parts, " ")
}

func main() {
	examples := []string{
		"привет мир",
		"hello WORLD",
		"ДоПустИм ЗабылОТключить cAPS",
		"",
	}

	for _, e := range examples {
		fmt.Printf("Исход: %q -> Результат: %q\n", e, capitalizeWords(e))
	}
}
