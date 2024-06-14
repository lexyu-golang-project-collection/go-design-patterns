package main

import "fmt"

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

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

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
