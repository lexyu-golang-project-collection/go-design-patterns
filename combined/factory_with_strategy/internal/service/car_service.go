package service

import (
	"fmt"

	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/factory"
	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/model"
	custom_type "github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/types"
)

type CarService struct{}

func NewCarService() *CarService {
	return &CarService{}
}

func (cs *CarService) BuildCar(carType custom_type.CarType) (*model.Car, error) {
	strategy := factory.GetCarBuildStrategy(carType)
	if strategy == nil {
		return nil, fmt.Errorf("failed strategy: %s", carType)
	}

	return strategy.BuildCar(), nil
}
