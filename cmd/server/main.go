package main

import (
	"github.com/Krynegal/smpl_calc_api.git/configs"
	"github.com/Krynegal/smpl_calc_api.git/internal/handlers"
	"log"
	"net/http"
)

func main() {
	cfg := configs.Get()
	log.Printf("Start server on port: %s", cfg.ServerPort)
	r := handlers.NewHandler().Router
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
