package main

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
	Build() *Pizza
}

type ConcretePizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *ConcretePizzaBuilder {
	return &ConcretePizzaBuilder{pizza: &Pizza{}}

}

func (cpb *ConcretePizzaBuilder) Build() *Pizza {
	return cpb.pizza
}

func main() {

}
