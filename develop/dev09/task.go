package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Парсинг аргументов командной строки
	url := flag.String("url", "", "URL сайта для загрузки")
	output := flag.String("output", "index.html", "имя файла для сохранения")
	flag.Parse()

	// Проверка наличия обязательного аргумента URL
	if *url == "" {
		fmt.Println("Необходимо указать URL сайта для загрузки")
		os.Exit(1)
	}

	// Выполнение HTTP-запроса
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println("Ошибка при выполнении HTTP-запроса:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP-запрос вернул статус: %s\n", resp.Status)
		os.Exit(1)
	}

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		os.Exit(1)
	}

	// Проверка, является ли путь к файлу директорией
	fileInfo, err := os.Stat(*output)
	if err == nil && fileInfo.IsDir() {
		// Если путь к файлу является директорией, добавляем к нему имя файла по умолчанию
		if !strings.HasSuffix(*output, "/") {
			*output += "/"
		}
		*output += "index.html"
	}

	// Запись данных в файл
	err = ioutil.WriteFile(*output, body, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		os.Exit(1)
	}

	fmt.Printf("Сайт успешно загружен в файл: %s\n", *output)
}
