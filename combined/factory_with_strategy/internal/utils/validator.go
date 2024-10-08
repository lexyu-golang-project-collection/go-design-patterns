package utils

import (
	"fmt"
	"strings"

	custom_type "github.com/lexyu-golang-project-collection/go-design-patterns/combined/factory_with_strategy/types"
)

func ValidatorCarType(carType string) error {
	fmt.Println("正在驗證車型: ", carType)

	validTypes := []string{string(custom_type.SportCar), string(custom_type.FamilyCar)}
	carType = strings.ToLower(carType)

	for _, vaild := range validTypes {
		if carType == vaild {
			fmt.Println("車型驗證通過")
			return nil
		}
	}

	return fmt.Errorf("無效的車型: %s", carType)
}
