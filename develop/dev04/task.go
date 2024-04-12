package main

import (
	"sort"
	"strings"
)

func findAnagrams(words *[]string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range *words {
		// Приведение слова к нижнему регистру и сортировка его символов
		word = strings.ToLower(word)
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		sortedWord := string(runes)

		// Добавление отсортированного слова в мапу анаграмм
		if _, ok := anagrams[sortedWord]; !ok {
			anagrams[sortedWord] = make([]string, 0)
		}
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	// Удаление множеств из одного элемента
	for key, value := range anagrams {
		if len(value) <= 1 {
			delete(anagrams, key)
		}
	}

	return anagrams
}

func main() {
	// Пример использования функции
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(&words)
	for key, value := range anagrams {
		println("Множество анаграмм для", key, ": ", strings.Join(value, ", "))
	}
}
