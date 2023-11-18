package main

import (
	"fmt"

	gu "github.com/lexyu-golang-project-collection/creational-patterns/factory/gun"
)

func main() {
	ak47, _ := gu.GetGun("ak47")
	musket, _ := gu.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g gu.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
