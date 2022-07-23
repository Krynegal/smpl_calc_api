package handlers

type response struct {
	Success bool
	Value   int64
}

type errResponse struct {
	Success bool
	Error   string
}
