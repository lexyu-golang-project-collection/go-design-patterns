package handler

import (
	"net/http"

	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/service"
	custom_type "github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/types"
)

type CarHandler struct {
	carService *service.CarService
}

func NewHandler(cs *service.CarService) *CarHandler {
	return &CarHandler{carService: cs}
}

func (handler *CarHandler) BuildCar(w http.ResponseWriter, r *http.Request) {

}

func (handler *CarHandler) BuildSportCar(w http.ResponseWriter, r *http.Request) {
	car, err := handler.carService.BuildCar(custom_type.SportCar)
}

func (handler *CarHandler) BuildFamilyCar(w http.ResponseWriter, r *http.Request) {

}
