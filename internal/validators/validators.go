package validators

import (
	"errors"
	"net/http"
	"strconv"
)

var (
	ParameterMissed = errors.New("parameter has been missed")
)

func validateParameters(o string) (float64, error) {
	operand, err := strconv.ParseFloat(o, 64)
	if err != nil {
		return 0, err
	}
	return operand, nil
}

func ParseParameters(r *http.Request) (float64, float64, error) {
	a, errA := validateParameters(r.URL.Query().Get("a"))
	if errA != nil {
		return 0, 0, ParameterMissed
	}
	b, errB := validateParameters(r.URL.Query().Get("b"))
	if errB != nil {
		return 0, 0, ParameterMissed
	}
	return a, b, nil
}
