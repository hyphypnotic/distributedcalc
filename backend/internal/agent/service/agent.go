package service

import (
	"fmt"
	"strconv"
	"sync"
)

const (
	minus          = '-'
	openBracket    = '('
	closeBracket   = ')'
	plus           = '+'
	division       = '/'
	multiplication = '*'
)

var priority = map[rune]int{
	openBracket:    0,
	closeBracket:   0,
	plus:           2,
	minus:          2,
	multiplication: 3,
	division:       3,
}

func setOrderOfOperations(expression string) string {
	stack := make([]rune, 0)
	output := make([]rune, 0)

	for _, char := range expression {
		switch char {
		case plus, minus, multiplication, division:
			for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[char] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		case openBracket:
			stack = append(stack, char)
		case closeBracket:
			for len(stack) > 0 && stack[len(stack)-1] != openBracket {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // Remove the open bracket
		default:
			output = append(output, char)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(output)
}

func EvaluateExpression(expression string) (float64, error) {
	expression = setOrderOfOperations(expression)
	stack := make([]float64, 0)
	var wg sync.WaitGroup
	mtx := &sync.Mutex{}

	for _, char := range expression {
		switch char {
		case '+', '-', '*', '/':
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression: not enough operands for operation")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			wg.Add(1)
			go func(op rune, op1, op2 float64) {
				defer wg.Done()
				var result float64
				switch op {
				case '+':
					result = op1 + op2
				case '-':
					result = op1 - op2
				case '*':
					result = op1 * op2
				case '/':
					if op2 == 0 {

						return
					}
					result = op1 / op2
				}

				mtx.Lock()
				stack = append(stack, result)
				mtx.Unlock()
			}(char, operand1, operand2)
		default:
			num, err := strconv.ParseFloat(string(char), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid operand: %s", string(char))
			}
			stack = append(stack, num)
		}
	}

	wg.Wait()

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression: unexpected number of operands")
	}

	return stack[0], nil
}
