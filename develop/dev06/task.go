package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func cut(fields []int, delimiter string, onlySeparated bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Разделение строки на колонки
		columns := strings.Split(line, delimiter)

		// Проверка, что строка содержит разделитель
		if len(columns) <= 1 && onlySeparated {
			continue
		}

		// Вывод выбранных колонок
		var selectedColumns []string
		for _, field := range fields {
			if field <= len(columns) {
				selectedColumns = append(selectedColumns, columns[field-1])
			}
		}
		fmt.Println(strings.Join(selectedColumns, delimiter))
	}
}

func main() {
	// Парсинг флагов командной строки
	var fieldsStr string
	flag.StringVar(&fieldsStr, "f", "", "выбрать поля (колонки)")
	var delimiter string
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	var onlySeparated bool
	flag.BoolVar(&onlySeparated, "s", false, "только строки с разделителем")
	flag.Parse()

	// Разбор выбранных колонок
	fields := make([]int, 0)
	for _, f := range strings.Split(fieldsStr, ",") {
		field := strings.TrimSpace(f)
		if field != "" {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: некорректное значение поля: %s\n", field)
				os.Exit(1)
			}
			fields = append(fields, num)
		}
	}

	// Проверка наличия выбранных колонок
	if len(fields) == 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: необходимо указать выбранные поля")
		os.Exit(1)
	}

	// Вызов функции cut с переданными флагами
	cut(fields, delimiter, onlySeparated)
}
