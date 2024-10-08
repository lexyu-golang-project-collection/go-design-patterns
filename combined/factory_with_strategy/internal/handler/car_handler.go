package handler

import (
	"net/http"

	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/service"
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

}

func (handler *CarHandler) BuildFamilyCar(w http.ResponseWriter, r *http.Request) {

}
