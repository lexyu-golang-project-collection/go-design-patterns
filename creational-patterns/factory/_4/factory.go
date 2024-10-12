package main

import (
	"errors"
	"fmt"
)

type ProductType int

const (
	BASIC ProductType = iota
	PREMIUM
)

// Product interface
type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
}

// ProductCreator interface
type ProductCreator interface {
	CreateProduct() Product
}

// BasicProduct struct
type BasicProduct struct {
	name        string
	price       float64
	description string
}

func (p *BasicProduct) GetName() string {
	return p.name
}

func (p *BasicProduct) GetPrice() float64 {
	return p.price
}

func (p *BasicProduct) GetDescription() string {
	return p.description
}

// BasicProductCreator struct
type BasicProductCreator struct{}

func (c *BasicProductCreator) CreateProduct() Product {
	return &BasicProduct{name: "基本產品", price: 100.0, description: "一個簡單的產品"}
}

// PremiumProduct struct
type PremiumProduct struct {
	BasicProduct
	features []string
}

func (p *PremiumProduct) GetDescription() string {
	return fmt.Sprintf("%s\n特色功能: %v", p.description, p.features)
}

// PremiumProductCreator struct
type PremiumProductCreator struct{}

func (c *PremiumProductCreator) CreateProduct() Product {
	return &PremiumProduct{
		BasicProduct: BasicProduct{name: "高級產品", price: 200.0, description: "一個更複雜的產品"},
		features:     []string{"特色功能1", "特色功能2"},
	}
}

// ProductFactory struct
type ProductFactory struct {
	creators map[ProductType]ProductCreator
}

// NewProductFactory creates a new ProductFactory with injected dependencies
func NewProductFactory(creators map[ProductType]ProductCreator) *ProductFactory {
	return &ProductFactory{creators: creators}
}

func (f *ProductFactory) CreateProduct(productType ProductType) (Product, error) {
	creator, exists := f.creators[productType]
	if !exists {
		return nil, errors.New("未知的產品類型")
	}
	return creator.CreateProduct(), nil
}

func main() {
	creators := map[ProductType]ProductCreator{
		BASIC:   &BasicProductCreator{},
		PREMIUM: &PremiumProductCreator{},
	}

	// 創建工廠實例，注入依賴
	factory := NewProductFactory(creators)

	// 使用工廠創建產品
	basicProduct, _ := factory.CreateProduct(BASIC)
	premiumProduct, _ := factory.CreateProduct(PREMIUM)

	fmt.Printf("基本產品: %s, 價格: %.2f\n", basicProduct.GetName(), basicProduct.GetPrice())
	fmt.Printf("高級產品: %s\n%s\n", premiumProduct.GetName(), premiumProduct.GetDescription())
}
