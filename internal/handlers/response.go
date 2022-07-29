package handlers

type Response struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error"`
}
