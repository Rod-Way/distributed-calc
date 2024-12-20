package calculator

import (
	"fmt"
	"strconv"
	"strings"

	m "distributedCalculator/pkg/calculator/models"
)

func Calculate(expression string) (float64, error) {

	if r, err := parseExpression(expression); err != nil {
		return 0, err
	} else {
		return r.Eval(), nil
	}
}

func parseExpression(expression string) (m.Node, error) {
	if tokens, err := tokenize(expression); err != nil {
		return nil, err
	} else {
		return buildTree(tokens)
	}
}

func tokenize(expression string) ([]string, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	var tokens []string
	currentNumber := ""

	for _, char := range expression {
		strChar := string(char)
		if strings.ContainsAny(strChar, "0123456789.") {
			currentNumber += strChar
		} else {
			if currentNumber != "" {
				tokens = append(tokens, currentNumber)
				currentNumber = ""
			}
			if isOperator(strChar) || strChar == "(" || strChar == ")" {
				tokens = append(tokens, strChar)
			} else {
				return nil, fmt.Errorf("invalid token: %s", strChar)
			}
		}
	}

	if currentNumber != "" {
		tokens = append(tokens, currentNumber)
	}

	return tokens, nil
}

func isOperator(s string) bool {
	for _, op := range operators {
		if op.Symbol == s {
			return true
		}
	}
	return false
}

func getOperator(symbol string) (m.Operator, error) {
	for _, op := range operators {
		if op.Symbol == symbol {
			return op, nil
		}
	}
	return m.Operator{}, fmt.Errorf("invalid operator: %s", symbol)
}

func buildTree(tokens []string) (m.Node, error) {
	var operandStack []m.Node
	var operatorStack []m.Operator

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			operandStack = append(operandStack, m.Operand{Value: num})
		} else if token == "(" {
			operatorStack = append(operatorStack, m.Operator{Symbol: token})
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Symbol != "(" {
				if err := applyOperator(&operandStack, &operatorStack); err != nil {
					return nil, err
				}
			}
			if len(operatorStack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else {
			op, err := getOperator(token)
			if err != nil {
				return nil, err
			}

			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Precedence >= op.Precedence {
				if err := applyOperator(&operandStack, &operatorStack); err != nil {
					return nil, err
				}
			}
			operatorStack = append(operatorStack, op)
		}
	}

	for len(operatorStack) > 0 {
		if err := applyOperator(&operandStack, &operatorStack); err != nil {
			return nil, err
		}
	}

	if len(operandStack) != 1 {
		return nil, fmt.Errorf("invalid expression")
	}

	return operandStack[0], nil
}

func applyOperator(operandStack *[]m.Node, operatorStack *[]m.Operator) error {
	if len(*operandStack) < 2 {
		return fmt.Errorf("not enough operands")
	}
	right := (*operandStack)[len(*operandStack)-1]
	*operandStack = (*operandStack)[:len(*operandStack)-1]
	left := (*operandStack)[len(*operandStack)-1]
	*operandStack = (*operandStack)[:len(*operandStack)-1]

	op := (*operatorStack)[len(*operatorStack)-1]
	*operatorStack = (*operatorStack)[:len(*operatorStack)-1]

	*operandStack = append(*operandStack, m.BinOp{Left: left, Right: right, Op: op})
	return nil
}
