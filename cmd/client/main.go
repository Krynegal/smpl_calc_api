package main

import (
	"bytes"
	"encoding/json"
	"github.com/Krynegal/smpl_calc_api.git"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	body := smpl_calc_api.Data{
		Operand1: smpl_calc_api.Operand{
			Value: "FF",
			Base:  16,
		},
		Operand2: smpl_calc_api.Operand{
			Value: "FF",
			Base:  1,
		},
		ToBase: 10,
	}
	r, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	log.Println(string(r))
	reqBody := bytes.NewBuffer(r)
	log.Println(reqBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/add", reqBody)
	if err != nil {
		panic("bad request")
	}

	resp, err := client.Do(req)
	if err != nil {
		panic("oops... there are some problems with server")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	log.Printf("body: %s", string(respBody))
}
