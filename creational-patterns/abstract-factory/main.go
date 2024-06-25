package main

import "fmt"

var (
	factories = map[string]func() ISportsFactory{
		"adidas": func() ISportsFactory { return &Adidas{} },
		"nike":   func() ISportsFactory { return &Nike{} },
	}
)

// interface Factory
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if factory, ok := factories[brand]; ok {
		return factory(), nil
	}
	return nil, fmt.Errorf("wrong brand type")
}

// Concrete Factory
type Adidas struct {
}

func (adidas *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (adidas *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct {
}

func (nike *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 16,
		},
	}
}

func (nike *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 16,
		},
	}
}

// Abstract class
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
	isShoe() ShoeDetails
}

type ShoeDetails struct{}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

func (s *Shoe) isShoe() ShoeDetails { return ShoeDetails{} }

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
	isShirt() ShirtDetails
}

type ShirtDetails struct{}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

func (s *Shirt) isShirt() ShirtDetails { return ShirtDetails{} }

// Concrete class
type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	fmt.Println("================")

	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)

	fmt.Println("================")

	// printShirtDetails(adidasShoe)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
