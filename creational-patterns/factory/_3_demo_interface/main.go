package main

import "fmt"

// Composition
type Factory interface {
	CreateProduct() Product
}

// Base
type Product interface {
	GetName() string
}

// Concrete Class
type ConcreteProduct struct {
	name string
}

func (p *ConcreteProduct) GetName() string {
	return p.name
}

// Concrete Factory
type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{name: "Concrete Product"}
}

func main() {
	factory := &ConcreteFactory{}
	product := factory.CreateProduct()
	fmt.Println(product.GetName())
}
