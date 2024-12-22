package handlers

import (
	m "distributedCalculator/internal/models"
	"distributedCalculator/pkg/calculator"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) PostCalculate(c echo.Context) error {
	expr := new(m.CalculateRequest)
	if err := c.Bind(&expr); err != nil {
		return c.JSON(http.StatusInternalServerError, m.CalculateErrorResponse{Error: "Internal server error"})
	}

	res, err := calculator.Calculate(expr.Expression)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, m.CalculateErrorResponse{Error: "Expression is not valid"})
	}

	result := fmt.Sprintf("%f", res)
	return c.JSON(http.StatusOK, m.CalculateSuccesfulResponse{Result: result})
}
