package main

import (
	"github.com/Krynegal/smpl_calc_api.git/configs"
	"github.com/Krynegal/smpl_calc_api.git/internal/handlers"
	"log"
	"net/http"
)

// @title Simple Calculator API
// @version 1.0
// @description Description

// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /

func main() {
	cfg := configs.Get()
	log.Printf("Start server on port: %s", cfg.ServerPort)
	r := handlers.NewHandler().Router
	log.Fatal(http.ListenAndServe(":8080", r))
}
