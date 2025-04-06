package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		input, operation string
		first, second    float64
	)

	fmt.Println("\n\tДобро пожаловать в программу Калькулятор!\n\tПоддерживаемые арифметические операторы: \nCложение: +\tВычитание: -\tУмножение: *\tДеление: /")
	fmt.Println("В качестве разделителя дробной части используйте знак точка: \".\"\n ")

	for i := 1; i < 3; i++ {

		fmt.Printf("Введите %d-ое число: ", i)
		fmt.Scan(&input)

		for _, err := strconv.ParseFloat(input, 64); err != nil; _, err = strconv.ParseFloat(input, 64) {
			fmt.Printf("Введенное вами значение не является числом! Пожалуйста введите %d-ое число: ", i)
			fmt.Scan(&input)
		}

		if i == 1 {
			first, _ = strconv.ParseFloat(input, 64)
		} else {
			second, _ = strconv.ParseFloat(input, 64)
			break
		}

		fmt.Println("Введите арифметический оператор: ")

	caseLoop:
		for {
			fmt.Scan(&input)
			if len(input) == 1 {
				switch rune(input[0]) {
				case 42:
					operation = input
					break caseLoop
				case 43:
					operation = input
					break caseLoop
				case 45:
					operation = input
					break caseLoop
				case 47:
					operation = input
					break caseLoop
				default:
					fmt.Println("Введенное вами значение не является доступным оператором! Пожалуйста введите арифметический оператор: ")
				}
			} else {
				fmt.Println("Введенное вами значение не является доступным оператором! Пожалуйста введите арифметический оператор: ")
			}
		}
	}

	switch rune(operation[0]) {
	case 43:
		fmt.Println("Результат сложения:", first+second)
	case 45:
		fmt.Println("Результат вычитания:", first-second)
	case 42:
		fmt.Println("Результат умножения:", first*second)
	case 47:
		fmt.Println("Результат деления:", first/second)
	}
}
