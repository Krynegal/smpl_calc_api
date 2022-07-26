package handlers

type Response struct {
	Success bool  `json:"success"`
	Result  int64 `json:"result"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error"`
}
