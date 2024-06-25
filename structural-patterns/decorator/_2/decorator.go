package main

import "fmt"

// Component Interface
type IPizza interface {
	getPrice() int
}

// Concrete Component
type VeggieMania struct {
}

func (v *VeggieMania) getPrice() int {
	return 20
}

// Concrete Decorator
type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 17
}

// Concrete Decorator
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 25
}

func main() {
	pizza := &VeggieMania{}
	fmt.Println("Pizza:", pizza.getPrice())

	pizza_with_cheese := &CheeseTopping{
		pizza: pizza,
	}
	fmt.Println("Pizza with Cheese:", pizza_with_cheese.getPrice())

	pizza_with_cheese_and_tomato := &TomatoTopping{
		pizza: pizza_with_cheese,
	}
	fmt.Println("Pizza with Cheese and Tomato:", pizza_with_cheese_and_tomato.getPrice())
}
