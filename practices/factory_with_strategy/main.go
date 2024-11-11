package main

import (
	"errors"
	"fmt"
)

// Operation types using iota
const (
	Add = iota
	Subtract
	Multiply
	Divide
	Mod
)

// Strategy interface with error handling
type Operation interface {
	Execute(a, b int) (int, error)
}

// Concrete strategies
type Addition struct{}

func (a *Addition) Execute(x, b int) (int, error) {
	return x + b, nil
}

type Subtraction struct{}

func (s *Subtraction) Execute(x, b int) (int, error) {
	return x - b, nil
}

type Multiplication struct{}

func (m *Multiplication) Execute(x, b int) (int, error) {
	return x * b, nil
}

type Division struct{}

func (d *Division) Execute(x, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除數不能為零 / division by zero")
	}
	return x / b, nil
}

type Modulus struct{}

func (m *Modulus) Execute(x, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("取模運算的除數不能為零 / modulus by zero")
	}
	return x % b, nil
}

// Factory method
func CreateOperation(operationType int) Operation {
	switch operationType {
	case Add:
		return &Addition{}
	case Subtract:
		return &Subtraction{}
	case Multiply:
		return &Multiplication{}
	case Divide:
		return &Division{}
	case Mod:
		return &Modulus{}
	default:
		return nil
	}
}

func main() {
	// Test the factory with different strategies
	numbers := []int{11, -3}
	operations := []int{Add, Subtract, Multiply, Divide, Mod}
	operationNames := []string{"加法/Addition", "減法/Subtraction",
		"乘法/Multiplication", "除法/Division",
		"取模/Modulus"}

	for i, op := range operations {
		operation := CreateOperation(op)
		result, err := operation.Execute(numbers[0], numbers[1])
		if err != nil {
			fmt.Printf("%s: Error - %v\n", operationNames[i], err)
			continue
		}
		fmt.Printf("%s: %d\n", operationNames[i], result)
	}

	// Test division by zero
	fmt.Println("\n測試除以零的情況 / Testing division by zero:")
	divOp := CreateOperation(Divide)
	if _, err := divOp.Execute(10, 0); err != nil {
		fmt.Printf("除法錯誤/Division Error: %v\n", err)
	}
}
