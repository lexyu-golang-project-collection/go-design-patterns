package main

import (
	"fmt"
)

// Factory
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

// Base
type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

// Concrete Ak47 : Composition
type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Concrete Musket : Composition
type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func main() {

	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

func printDetails(ig IGun) {
	fmt.Printf("Gun: %s", ig.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", ig.GetPower())
	fmt.Println()
}
