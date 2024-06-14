package main

import (
	"fmt"
)

func main() {

	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(ig IGun) {
	fmt.Printf("Gun: %s", ig.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", ig.GetPower())
	fmt.Println()
}
