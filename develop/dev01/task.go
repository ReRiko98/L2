package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем точное время от сервера NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Выводим ошибку в STDERR
		fmt.Fprintf(os.Stderr, "Ошибка получения времени: %v\n", err)
		// Возвращаем ненулевой код выхода в OS
		os.Exit(1)
	}

	// Печатаем точное время
	fmt.Println("Точное время (NTP):", ntpTime)
}
