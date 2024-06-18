package main

import (
	"fmt"
)

func main() {
	calculator()
}

func calculator() {
	var num1, num2 float64

	fmt.Print("Введите первое число: ")

	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Printf("Ошибка ввода первого числа: %v", err)
		return
	}

	fmt.Print("Введите арифметический оператор (+, -, *, /): ")

	var operator string
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

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("На ноль делить нельзя")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Неизвестный арифметический оператор")
	}
}
