package main

import (
	"fmt"
	"log"
	"net/http"

	_ "backend/docs"
	swagger "github.com/swaggo/http-swagger/v2"

	"backend/config"
	"backend/internal/database"
	"backend/pkg/middleware"

	desenvolvedorRepository "backend/internal/repository/desenvolvedor"
	nivelRepository "backend/internal/repository/nivel"

	desenvolvedorUseCase "backend/internal/usecase/desenvolvedor"
	nivelUseCase "backend/internal/usecase/nivel"
)

var (
	cfg = config.NewConfig().WithMySQL()

	mysql = database.NewMySQLDatabase(cfg)

	dr = desenvolvedorRepository.NewMySQLRepository(mysql)
	nr = nivelRepository.NewMySQLRepository(mysql)

	du = desenvolvedorUseCase.NewUseCase(dr)
	nu = nivelUseCase.NewUseCase(nr)
)

// @title API Desafio Fullstack Gazin
// @description API desenvolvida para o desafio fullstack Gazin
// @version 1.0
func main() {
	// API Routes
	routes := http.NewServeMux()
	// Niveis
	routes.HandleFunc("GET /niveis", nu.Index)
	routes.HandleFunc("GET /niveis/select-options", nu.IndexForSelectOptions)
	routes.HandleFunc("POST /niveis", nu.Create)
	routes.HandleFunc("PUT /niveis/{id}", nu.Update)
	routes.HandleFunc("DELETE /niveis/{id}", nu.Delete)
	// Desenvolvedores
	routes.HandleFunc("GET /desenvolvedores", du.Index)
	routes.HandleFunc("POST /desenvolvedores", du.Create)
	routes.HandleFunc("PUT /desenvolvedores/{id}", du.Update)
	routes.HandleFunc("DELETE /desenvolvedores/{id}", du.Delete)

	// Server Multiplexer
	api := http.NewServeMux()
	api.Handle("/", http.DefaultServeMux)
	api.Handle("/api/", http.StripPrefix("/api", routes))
	api.Handle("/docs/", swagger.Handler())

	// Middleware Stack
	cors := middleware.NewCORS(cfg)
	middlewares := middleware.Chain(
		middleware.Logger,
		cors.Handler,
		middleware.Json,
	)

	// Server Configurations
	server := http.Server{
		Addr:    cfg.ServerPort,
		Handler: middlewares(api),
	}

	// Server Start
	log.Println("Server running on", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
