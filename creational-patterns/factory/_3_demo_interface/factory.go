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

// Product Base
type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	Update(name string, price float64, description string) error
}

// Factory
type Factory interface {
	CreateProduct(productType ProductType) (Product, error)
	CreateBulkProducts(productType ProductType, count int) ([]Product, error)
	GetProductInfo(productType ProductType) string
}

// Concrete Product
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

func (p *BasicProduct) Update(name string, price float64, description string) error {
	p.name = name
	p.price = price
	p.description = description
	return nil
}

// Concrete Product
type PremiumProduct struct {
	BasicProduct
	features []string
}

func (p *PremiumProduct) GetDescription() string {
	return fmt.Sprintf("%s\n特色功能: %v", p.description, p.features)
}

// Concrete Factory
type ProductFactory struct{}

func (f *ProductFactory) CreateProduct(productType ProductType) (Product, error) {
	switch productType {
	case BASIC:
		return &BasicProduct{name: "基本產品", price: 100.0, description: "一個簡單的產品"}, nil
	case PREMIUM:
		return &PremiumProduct{
			BasicProduct: BasicProduct{name: "高級產品", price: 200.0, description: "一個更複雜的產品"},
			features:     []string{"特色功能1", "特色功能2"},
		}, nil
	default:
		return nil, errors.New("未知的產品類型")
	}
}

func (f *ProductFactory) CreateBulkProducts(productType ProductType, count int) ([]Product, error) {
	products := make([]Product, count)
	var err error
	for i := 0; i < count; i++ {
		products[i], err = f.CreateProduct(productType)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func (f *ProductFactory) GetProductInfo(productType ProductType) string {
	switch productType {
	case BASIC:
		return "基本產品: 一個簡單、基礎的產品"
	case PREMIUM:
		return "高級產品: 一個更先進的產品，具有額外的特色功能"
	default:
		return "未知的產品類型"
	}
}

func main() {
	factory := &ProductFactory{}

	// 創建單個產品
	product, err := factory.CreateProduct(PREMIUM)
	if err != nil {
		fmt.Println("錯誤:", err)
		return
	}
	fmt.Println("產品:", product.GetName())
	fmt.Println("描述:", product.GetDescription())
	fmt.Println("價格:", product.GetPrice())

	// 批量創建產品
	bulkProducts, err := factory.CreateBulkProducts(BASIC, 3)
	if err != nil {
		fmt.Println("錯誤:", err)
		return
	}
	fmt.Println("\n批量產品:")
	for i, p := range bulkProducts {
		fmt.Printf("%d. %s, %s, %v\n", i+1, p.GetName(), p.GetDescription(), p.GetPrice())
	}

	// 獲取產品信息
	fmt.Println("\n產品信息:")
	fmt.Println(factory.GetProductInfo(PREMIUM))
}
