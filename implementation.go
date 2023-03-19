package lab2

import (
	"errors"
	"fmt"
	"strings"
)

// prefixToInfix converts prefix notation to infix notation.
func prefixToInfix(input string) (string, error) {
	var stack []string
	const operators = "+-*/^"
	const errorMsg = "invalid expression"

	if len(input) == 0 {
		return "", errors.New(errorMsg)
	}

	inputs := strings.Split(input, " ")
	var symbol, expression1, expression2 string

	for i := len(inputs) - 1; i >= 0; i-- {
		symbol = inputs[i]

		if !strings.Contains(operators, symbol) {
			stack = append(stack, symbol)
		} else {
			if len(stack) < 2 {
				return "", errors.New(errorMsg)
			}

			expression1, stack = pop(stack)
			expression2, stack = pop(stack)

			newExpression := join(expression1, expression2, symbol)
			stack = append(stack, newExpression)
		}
	}

	res, stack := pop(stack)
	if len(stack) > 0 {
		return "", errors.New(errorMsg)
	}

	return res, nil
}

// pop removes and returns the last element from the slice.
func pop(str []string) (string, []string) {
	if len(str) == 0 {
		return "", []string{}
	}

	return str[len(str)-1], str[:len(str)-1]
}

// join returns a new expression by concatenating expression1, symbol, and expression2.
// If either expression contains spaces and symbol is '*', '/', or '^', the expression is enclosed in parentheses.
func join(expression1, expression2, symbol string) string {
	if strings.Contains(expression1, " ") && strings.Contains("*/^", symbol) {
		expression1 = "(" + expression1 + ")"
	}
	if strings.Contains(expression2, " ") && strings.Contains("*/^", symbol) {
		expression2 = "(" + expression2 + ")"
	}

	return expression1 + " " + symbol + " " + expression2
}