package main

import (
	"bytes"
	"encoding/json"
	"github.com/Krynegal/smpl_calc_api.git/internal/types"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	operands := [2]types.Operand{
		{
			Value: "35",
			Base:  10,
		},
		{
			Value: "0",
			Base:  10,
		},
	}
	body := types.Data{
		Operands: operands,
		ToBase:   10,
	}
	r, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	log.Println(string(r))
	reqBody := bytes.NewBuffer(r)
	log.Println(reqBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/div", reqBody)
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
