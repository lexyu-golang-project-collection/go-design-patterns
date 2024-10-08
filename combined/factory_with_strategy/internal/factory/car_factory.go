package factory

import (
	"github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/internal/strategy"
	custom_type "github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/types"
)

func GetCarBuildStrategy(carType custom_type.CarType) strategy.CarBuildStrategy {
	switch carType {
	case custom_type.SportCar:
		return &strategy.SportCarStrategy{}
	case custom_type.FamilyCar:
		return &strategy.FamilyCarStrategy{}
	default:
		return nil
	}
}
