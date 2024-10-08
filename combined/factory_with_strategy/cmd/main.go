package main

import (
	"log"
	"net/http"

	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/handler"
	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/middleware"
	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/service"
	logger "github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/pkg"
)

func main() {
	carService := service.NewCarService()
	carHandler := handler.NewHandler(carService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/car", middleware.Logging(middleware.Auth(carHandler.BuildCar)))
	mux.HandleFunc("GET /api/car/sport", middleware.Logging(middleware.Auth(carHandler.BuildSportCar)))
	mux.HandleFunc("GET /api/car/family", middleware.Logging(middleware.Auth(carHandler.BuildFamilyCar)))

	logger.Info("Server running on http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", mux))
}
