package handlers

import (
	"encoding/json"
	"github.com/Krynegal/smpl_calc_api.git/internal/types"
	"github.com/Krynegal/smpl_calc_api.git/internal/validators"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool
	Err     string
	Value   int64
}

func ConvertInt(operand types.Operand) (int64, string, error) {
	i, err := strconv.ParseInt(operand.Value, operand.Base, 64)
	if err != nil {
		return 0, "", err
	}
	return i, strconv.FormatInt(i, 10), nil
}

func GetOperands(data types.Data) (int64, int64, error) {
	var op1, op2 int64
	op1, _, err := ConvertInt(data.Operands[0])
	op2, _, err = ConvertInt(data.Operands[1])
	if err != nil {
		return 0, 0, err
	}
	return op1, op2, nil
}

func HandlerAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := validators.GetData(r)
		op1, op2, err := GetOperands(data)
		sum := op1 + op2
		log.Printf("sum: %v", sum)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error()}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: sum}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerSub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := validators.GetData(r)
		op1, op2, err := GetOperands(data)
		sub := op1 - op2
		log.Printf("sub: %v", sub)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error()}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: sub}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerMul() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := validators.GetData(r)
		op1, op2, err := GetOperands(data)
		mul := op1 * op2
		log.Printf("mul: %v", mul)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error()}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: mul}
		json.NewEncoder(w).Encode(&resResponse)
	}
}

func HandlerDiv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := validators.GetData(r)
		op1, op2, err := GetOperands(data)
		if op2 == 0 {
			resResponse := Response{Success: false, Err: "divide by zero"}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		div := op1 / op2
		log.Printf("div: %v", div)
		if err != nil {
			resResponse := Response{Success: false, Err: err.Error()}
			json.NewEncoder(w).Encode(&resResponse)
			return
		}
		resResponse := Response{Success: true, Err: "", Value: div}
		json.NewEncoder(w).Encode(&resResponse)
	}
}
