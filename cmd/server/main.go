package main

import (
	"fmt"
	"github.com/Krynegal/smpl_calc_api.git/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Start server")
	r.HandleFunc("/api/add", handlers.HandlerAdd())
	r.HandleFunc("/api/sub", handlers.HandlerSub())
	r.HandleFunc("/api/mul", handlers.HandlerMul())
	r.HandleFunc("/api/div", handlers.HandlerDiv())
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
