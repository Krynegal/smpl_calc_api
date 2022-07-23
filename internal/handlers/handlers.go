package handlers

import (
	"encoding/json"
	"github.com/Krynegal/smpl_calc_api.git"
	"github.com/Krynegal/smpl_calc_api.git/internal/validators"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Router.HandleFunc("/api/add", h.HandlerAdd).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/sub", h.HandlerSub).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/mul", h.HandlerMul).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/div", h.HandlerDiv).Methods(http.MethodPost)
	return h
}

func ConvertInt(operand smpl_calc_api.Operand) (int64, string, error) {
	i, err := strconv.ParseInt(operand.Value, operand.Base, 64)
	if err != nil {
		return 0, "", err
	}
	return i, strconv.FormatInt(i, 10), nil
}

func GetOperands(data smpl_calc_api.Data) (int64, int64, error) {
	var op1, op2 int64
	op1, _, err := ConvertInt(data.Operand1)
	op2, _, err = ConvertInt(data.Operand2)
	if err != nil {
		return 0, 0, err
	}
	return op1, op2, nil
}

func (h *Handler) HandlerAdd(w http.ResponseWriter, r *http.Request) {
	data, err := validators.GetData(r)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	if err = data.CheckData(); err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	op1, op2, err := GetOperands(data)
	sum := op1 + op2
	log.Printf("sum: %v", sum)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	resResponse := response{Success: true, Value: sum}
	json.NewEncoder(w).Encode(&resResponse)
}

func (h *Handler) HandlerSub(w http.ResponseWriter, r *http.Request) {
	data, err := validators.GetData(r)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	if err = data.CheckData(); err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	op1, op2, err := GetOperands(data)
	sub := op1 - op2
	log.Printf("sub: %v", sub)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	resResponse := response{Success: true, Value: sub}
	json.NewEncoder(w).Encode(&resResponse)
}

func (h *Handler) HandlerMul(w http.ResponseWriter, r *http.Request) {
	data, err := validators.GetData(r)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	if err = data.CheckData(); err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	op1, op2, err := GetOperands(data)
	mul := op1 * op2
	log.Printf("mul: %v", mul)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	resResponse := response{Success: true, Value: mul}
	json.NewEncoder(w).Encode(&resResponse)
}

func (h *Handler) HandlerDiv(w http.ResponseWriter, r *http.Request) {
	data, err := validators.GetData(r)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	if err = data.CheckData(); err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	op1, op2, err := GetOperands(data)
	if op2 == 0 {
		resResponse := errResponse{Success: false, Error: "divide by zero"}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	div := op1 / op2
	log.Printf("div: %v", div)
	if err != nil {
		resResponse := errResponse{Success: false, Error: err.Error()}
		json.NewEncoder(w).Encode(&resResponse)
		return
	}

	resResponse := response{Success: true, Value: div}
	json.NewEncoder(w).Encode(&resResponse)
}
