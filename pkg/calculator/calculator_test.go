package calculator_test

import (
	"testing"

	c "distributedCalculator/pkg/calculator"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "(())",
			expression:     "2+(2*(1+2))",
			expectedResult: 8.0,
		},

		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := c.Calculate(testCase.expression)
			if err != nil {
				t.Errorf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Errorf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}
}
