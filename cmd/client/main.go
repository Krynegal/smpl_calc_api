package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	resp, err := client.Get("http://localhost:8080/api/mul?a=13.5&b=")
	if err != nil {
		log.Println("can't get response")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	log.Printf("body: %s", string(respBody))
}
