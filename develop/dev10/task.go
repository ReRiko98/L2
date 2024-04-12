package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут на подключение")
	flag.Parse()

	// Проверка наличия аргументов хоста и порта
	if flag.NArg() != 2 {
		fmt.Println("Использование: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	// Установка таймера для завершения программы при истечении таймаута
	timer := time.NewTimer(*timeout)
	defer timer.Stop()

	// Подключение к серверу
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Создание канала для сигналов Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Запуск горутины для чтения сокета
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Ошибка при чтении сокета:", err)
				return
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	// Ожидание сигнала Ctrl+C или завершения таймера
	select {
	case <-signalChan:
		fmt.Println("\nПрограмма завершена по сигналу Ctrl+C")
	case <-timer.C:
		fmt.Println("Превышено время ожидания подключения")
	}
}
