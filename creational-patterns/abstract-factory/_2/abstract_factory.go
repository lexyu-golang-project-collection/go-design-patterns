package main

import (
	"fmt"
)

// Product Base
type Chair interface {
	Describe()
	GetStyle() string
}

type Table interface {
	Describe()
	GetStyle() string
}

// Abstract Factory
type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
}

// Concrete Product
type ModernChair struct{}

func (c *ModernChair) Describe() {
	fmt.Println("這是一張舒適的現代風格椅子，採用簡潔的線條設計")
}

func (c *ModernChair) GetStyle() string {
	return "現代"
}

type ModernTable struct{}

func (t *ModernTable) Describe() {
	fmt.Println("這是一張時尚的現代風格桌子，具有光滑的表面和簡約的外觀")
}

func (t *ModernTable) GetStyle() string {
	return "現代"
}

// Concrete Product
type ClassicChair struct{}

func (c *ClassicChair) Describe() {
	fmt.Println("這是一張優雅的古典風格椅子，帶有精緻的雕刻和軟墊")
}

func (c *ClassicChair) GetStyle() string {
	return "古典"
}

type ClassicTable struct{}

func (t *ClassicTable) Describe() {
	fmt.Println("這是一張華麗的古典風格桌子，具有豐富的裝飾細節和堅固的結構")
}

func (t *ClassicTable) GetStyle() string {
	return "古典"
}

// Concrete Factory * 2
type ModernFurnitureFactory struct{}

func (f *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (f *ModernFurnitureFactory) CreateTable() Table {
	return &ModernTable{}
}

type ClassicFurnitureFactory struct{}

func (f *ClassicFurnitureFactory) CreateChair() Chair {
	return &ClassicChair{}
}

func (f *ClassicFurnitureFactory) CreateTable() Table {
	return &ClassicTable{}
}

// Client
func CreateFurnitureSet(factory FurnitureFactory) {
	chair := factory.CreateChair()
	table := factory.CreateTable()

	fmt.Printf("創建了一套%s風格的傢俱\n", chair.GetStyle())
	chair.Describe()
	table.Describe()
}

func main() {
	modernFactory := &ModernFurnitureFactory{}
	classicFactory := &ClassicFurnitureFactory{}

	fmt.Println("使用現代傢俱工廠：")
	CreateFurnitureSet(modernFactory)

	fmt.Println("\n使用古典傢俱工廠：")
	CreateFurnitureSet(classicFactory)
}
