package handlers

import (
	"encoding/json"
	"github.com/Krynegal/smpl_calc_api.git/internal/validators"
	"log"
	"net/http"
)

type Response struct {
	Success bool
	Err     string
	Value   float64
}

func HandlerAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operand1, operand2, err := validators.ParseParameters(r)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error(), Value: operand1 + operand2}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: operand1 + operand2}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerSub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operand1, operand2, err := validators.ParseParameters(r)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error(), Value: operand1 - operand2}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: operand1 - operand2}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerMul() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operand1, operand2, err := validators.ParseParameters(r)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error(), Value: operand1 * operand2}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: operand1 * operand2}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerDiv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operand1, operand2, err := validators.ParseParameters(r)
		log.Println(operand1, operand2)
		if err != nil || operand1 == 0.0 || operand2 == 0.0 {
			resResponse := Response{Success: false, Err: "", Value: 0}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: operand1 / operand2}
		json.NewEncoder(w).Encode(&resResponse)
	}
}
