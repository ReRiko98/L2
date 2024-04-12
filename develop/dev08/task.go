package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}

		// Удаляем символ новой строки из ввода
		input = strings.TrimSpace(input)

		// Проверка на команду выхода
		if input == "\\quit" {
			break
		}

		// Разбиение ввода на команду и аргументы
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		// Обработка команды
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Необходимо указать аргумент для команды cd")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			if len(args) < 2 {
				fmt.Println("Необходимо указать аргумент для команды echo")
				continue
			}
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Необходимо указать PID для команды kill")
				continue
			}
			cmd := exec.Command("kill", args[1])
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды kill:", err)
			}
		case "ps":
			cmd := exec.Command("ps")
			out, err := cmd.Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды ps:", err)
				continue
			}
			fmt.Println(string(out))
		default:
			// Выполнение внешней команды
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
			}
		}
	}
}
