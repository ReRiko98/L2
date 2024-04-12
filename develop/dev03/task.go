package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	// Парсинг флагов командной строки
	column := flag.Int("k", -1, "номер колонки для сортировки (по умолчанию -1)")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	byMonth := flag.Bool("M", false, "сортировать по названию месяца")
	ignoreTrailingSpace := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "проверять отсортированы ли данные")
	numericWithSuffix := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	// Чтение строк из файла
	lines, err := readLines(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Применение сортировки в зависимости от флагов
	if *byMonth {
		sortByMonth(lines)
	} else if *numericWithSuffix {
		sortNumericWithSuffix(lines)
	} else {
		sortLines(lines, *column, *numeric, *reverse, *ignoreTrailingSpace)
	}

	// Удаление повторяющихся строк, если указан флаг -u
	if *unique {
		lines = uniqueLines(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}

	// Проверка отсортированности данных, если указан флаг -c
	if *checkSorted {
		if isSorted(lines, *reverse) {
			fmt.Println("Данные отсортированы.")
		} else {
			fmt.Println("Данные не отсортированы.")
		}
	}
}

// readLines читает строки из файла и возвращает их в виде среза строк.
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// sortLines сортирует строки в соответствии с заданными параметрами.
func sortLines(lines []string, column int, numeric, reverse, ignoreTrailingSpace bool) {
	sort.SliceStable(lines, func(i, j int) bool {
		a, b := lines[i], lines[j]

		if ignoreTrailingSpace {
			a = strings.TrimSpace(a)
			b = strings.TrimSpace(b)
		}

		if column != -1 {
			fieldsA := strings.Fields(a)
			fieldsB := strings.Fields(b)

			if len(fieldsA) <= column-1 || len(fieldsB) <= column-1 {
				return a < b
			}

			a = fieldsA[column-1]
			b = fieldsB[column-1]
		}

		if numeric {
			valA, errA := strconv.ParseFloat(a, 64)
			valB, errB := strconv.ParseFloat(b, 64)
			if errA == nil && errB == nil {
				return valA < valB
			}
		}

		if reverse {
			return a > b
		}
		return a < b
	})
}

// sortByMonth сортирует строки по названию месяца.
func sortByMonth(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		dateA, errA := time.Parse("Jan", lines[i])
		dateB, errB := time.Parse("Jan", lines[j])
		if errA != nil || errB != nil {
			return lines[i] < lines[j]
		}
		return dateA.Before(dateB)
	})
}

// sortNumericWithSuffix сортирует строки по числовому значению с учетом суффиксов.
func sortNumericWithSuffix(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		valA, suffixA := splitNumericWithSuffix(lines[i])
		valB, suffixB := splitNumericWithSuffix(lines[j])

		if valA == valB {
			return suffixA < suffixB
		}
		return valA < valB
	})
}

// splitNumericWithSuffix разделяет строку на числовое значение и суффикс.
// splitNumericWithSuffix разделяет строку на числовое значение и суффикс.
func splitNumericWithSuffix(s string) (int, string) {
	var val int
	var suffix string
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			val, _ = strconv.Atoi(s[:i+1])
			suffix = s[i+1:]
			break
		}
	}
	return val, suffix
}



// uniqueLines удаляет повторяющиеся строки из среза строк.
func uniqueLines(lines []string) []string {
	seen := make(map[string]struct{})
	unique := []string{}
	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			unique = append(unique, line)
		}
	}
	return unique
}

// isSorted проверяет, отсортированы ли данные.
func isSorted(lines []string, reverse bool) bool {
	if reverse {
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] < lines[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] > lines[i+1] {
				return false
			}
		}
	}
	return true
}
