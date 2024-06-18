package calculator

import (
	"errors"
	"fmt"
)

// calculator - неэкспортируемая структура
type calculator struct{}

// NewCalculator - экспортируемая функция-конструктор
// для создания экземпляра структуры calculator
func NewCalculator() *calculator {
	return &calculator{}
}

// Calculate - экспортируемый метод для выполнения арифметических операций
func (c *calculator) Calculate(num1, num2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, errors.New("неизвестный оператор")
	}
}

// add - неэкспортируемый метод для сложения
func (c *calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

// subtract - неэкспортируемый метод для вычитания
func (c *calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}

// multiply - неэкспортируемый метод для умножения
func (c *calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// divide - неэкспортируемый метод для деления
func (c *calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("Делитель не может быть нулем.")
		return 0
	}
	return num1 / num2
}
