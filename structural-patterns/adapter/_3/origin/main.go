package main

import (
	"encoding/xml"
	"fmt"
)

// XMLStockData represents the structure of our XML data
type XMLStockData struct {
	XMLName xml.Name `xml:"stock"`
	Symbol  string   `xml:"symbol"`
	Price   float64  `xml:"price"`
}

// StockDataProvider simulates the external XML data provider
type StockDataProvider struct{}

func (p *StockDataProvider) GetStockData() string {
	return `
		<stock>
			<symbol>AAPL</symbol>
			<price>150.25</price>
		</stock>
	`
}

// Application represents the core application that works with XML
type Application struct{}

func (a *Application) ProcessStockData(xmlData string) (*XMLStockData, error) {
	var stockData XMLStockData
	err := xml.Unmarshal([]byte(xmlData), &stockData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}
	return &stockData, nil
}

// Client uses the processed XML data
type Client struct{}

func (c *Client) UseStockData(stockData *XMLStockData) {
	fmt.Printf("Using stock data: Symbol=%s, Price=%.2f\n", stockData.Symbol, stockData.Price)
}

func main() {
	provider := &StockDataProvider{}
	app := &Application{}
	client := &Client{}

	// Get XML data from provider
	xmlData := provider.GetStockData()
	fmt.Println("Original XML data:")
	fmt.Println(xmlData)

	// Process XML data
	stockData, err := app.ProcessStockData(xmlData)
	if err != nil {
		fmt.Println("Error processing data:", err)
		return
	}

	// Client uses the XML data
	client.UseStockData(stockData)
}
