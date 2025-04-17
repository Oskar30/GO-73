package main

import (
	"fmt"
	calc "module5/pkg/calc"
)

func main() {
	var (
		opr  string
		x, y float64
	)

	fmt.Println("\n\tДобро пожаловать в программу Калькулятор!\n\tПоддерживаемые арифметические операторы: \nCложение: +\tВычитание: -\tУмножение: *\tДеление: /")
	fmt.Println("В качестве разделителя дробной части используйте знак точка: \".\"\n ")

	fmt.Print("Введите первое число: ")
	for _, err := fmt.Scan(&x); err != nil; _, err = fmt.Scan(&x) {
		fmt.Print("Введенное вами значение не является числом!\nПожалуйста введите первое число: ")
	}

	fmt.Print("\nВведите арифметический оператор: ")
	for _, err := fmt.Scan(&opr); err != nil || opr != "+" && opr != "-" && opr != "*" && opr != "/"; _, err = fmt.Scan(&opr) {
		fmt.Print("Введенное вами значение не является доступным оператором!\nПожалуйста введите арифметический оператор: ")
	}

	fmt.Print("\nВведите второе число: ")
	for _, err := fmt.Scan(&y); err != nil; _, err = fmt.Scan(&y) {
		fmt.Print("Введенное вами значение не является числом!\nПожалуйста введите второе число: ")
	}

	exemp := calc.New().Calculate(x, y, opr)

	fmt.Println("\nРезультат: ", exemp, "\n ")
}
