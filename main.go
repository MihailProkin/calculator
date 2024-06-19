package main

import (
	"fmt"

	"github.com/mihailprokin/calculator/calc"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")

	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Printf("Ошибка ввода первого числа: %v", err)
		return
	}

	fmt.Print("Введите арифметический оператор (+, -, *, /): ")

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Printf("Ошибка ввода арифметического оператора: %v", err)
		return
	}

	fmt.Print("Введите второе число: ")

	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Printf("Ошибка ввода второго числа: %v", err)
		return
	}

	calc := calc.NewCalculator()
	result, err := calc.Calculate(num1, num2, operator)
	if err != nil {
		fmt.Printf("Ошибка при выполнении арифметической операции: %v", err)
		return
	}

	fmt.Printf("Результат операции: %v", result)
}
