package calculator

import m "distributedCalculator/pkg/calculator/models"

var operators = []m.Operator{
	{"*", 2, multiply},
	{"/", 2, divide},
	{"+", 1, add},
	{"-", 1, subtract},
}

func multiply(l, r float64) float64 {
	return l * r
}

func divide(l, r float64) float64 {
	if r == 0 {
		panic("division by zero")
	}
	return l / r
}
func add(l, r float64) float64 {
	return l + r
}

func subtract(l, r float64) float64 {
	return l - r
}
