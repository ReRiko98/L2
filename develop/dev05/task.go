package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func grep(filename string, pattern string, after int, before int, context int, ignoreCase bool, invert bool, fixed bool, lineNumber bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		match := false

		if fixed {
			match = strings.Contains(line, pattern)
		} else {
			if ignoreCase {
				line = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
			}
			match = strings.Contains(line, pattern)
		}

		if (invert && !match) || (!invert && match) {
			lines = append(lines, line)
		}
	}

	printLines(lines, after, before, context, lineNumber)

	return nil
}

func printLines(lines []string, after int, before int, context int, lineNumber bool) {
	for i, line := range lines {
		if lineNumber {
			fmt.Printf("%d: ", i+1)
		}
		fmt.Println(line)

		if context > 0 {
			start := i - before
			end := i + after + 1

			if start < 0 {
				start = 0
			}
			if end > len(lines) {
				end = len(lines)
			}

			for j := start; j < end; j++ {
				if j == i {
					continue
				}
				fmt.Printf("\t%s\n", lines[j])
			}
		}
	}
}

func main() {
	// Парсинг флагов командной строки
	filename := flag.String("file", "", "имя файла для поиска")
	pattern := flag.String("pattern", "", "шаблон для поиска")
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNumber := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()

	// Проверка обязательных параметров
	if *filename == "" || *pattern == "" {
		fmt.Println("Имя файла и шаблон для поиска должны быть указаны")
		os.Exit(1)
	}

	// Вызов функции grep с переданными флагами
	err := grep(*filename, *pattern, *after, *before, *context, *ignoreCase, *invert, *fixed, *lineNumber)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
}
