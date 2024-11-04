package main

import "fmt"

type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}

type PizzaBuilder interface {
	setDough(string) PizzaBuilder
	setSauce(string) PizzaBuilder
	setCheese(string) PizzaBuilder
	setToppings([]string) PizzaBuilder
	build() *Pizza
}

type ConcretePizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *ConcretePizzaBuilder {
	return &ConcretePizzaBuilder{pizza: &Pizza{}}

}

func (cpb *ConcretePizzaBuilder) setDough(dough string) PizzaBuilder {
	cpb.pizza.dough = dough
	return cpb
}

func (cpb *ConcretePizzaBuilder) setSauce(sauce string) PizzaBuilder {
	cpb.pizza.sauce = sauce
	return cpb
}

func (cpb *ConcretePizzaBuilder) setCheese(cheese string) PizzaBuilder {
	cpb.pizza.cheese = cheese
	return cpb
}

func (cpb *ConcretePizzaBuilder) setToppings(toppings []string) PizzaBuilder {
	cpb.pizza.toppings = toppings
	return cpb
}

func (cpb *ConcretePizzaBuilder) build() *Pizza {
	return cpb.pizza
}

func main() {
	pizzaBuilder := NewPizzaBuilder()
	pizza := pizzaBuilder.
		setDough("Thin Crust").
		setSauce("Tomato").
		setCheese("Mozzarella").
		setToppings([]string{"Mushrooms", "Olives", "Onions"}).
		build()

	fmt.Println(pizza)
}
