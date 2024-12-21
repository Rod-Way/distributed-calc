package models

type (
	CalculateRequest struct {
		Expression string `json:"expression" form:"expression" query:"expression"`
	}
	CalculateSuccesfulResponse struct {
		Result string `json:"result"`
	}
	CalculateErrorResponse struct {
		Error string `json:"error"`
	}
)
