package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// XMLStockData represents the structure of our XML data
type XMLStockData struct {
	XMLName xml.Name `xml:"stock"`
	Symbol  string   `xml:"symbol"`
	Price   float64  `xml:"price"`
}

// JSONStockData represents the structure of our JSON data
type JSONStockData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
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

// Application represents the core application that initially works with XML
type Application struct{}

func (a *Application) ProcessStockData(xmlData string) (*XMLStockData, error) {
	var stockData XMLStockData
	err := xml.Unmarshal([]byte(xmlData), &stockData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}
	return &stockData, nil
}

// XMLToJSONAdapter is the adapter that converts XML to JSON
type XMLToJSONAdapter struct {
	app *Application
}

func (a *XMLToJSONAdapter) ConvertToJSON(xmlData string) (*JSONStockData, error) {
	xmlStock, err := a.app.ProcessStockData(xmlData)
	if err != nil {
		return nil, err
	}

	jsonStock := &JSONStockData{
		Symbol: xmlStock.Symbol,
		Price:  xmlStock.Price,
	}

	return jsonStock, nil
}

// Client can now use either XML or JSON data
type Client struct{}

func (c *Client) UseXMLData(stockData *XMLStockData) {
	fmt.Printf("Using XML stock data: Symbol=%s, Price=%.2f\n", stockData.Symbol, stockData.Price)
}

func (c *Client) UseJSONData(stockData *JSONStockData) {
	fmt.Printf("Using JSON stock data: Symbol=%s, Price=%.2f\n", stockData.Symbol, stockData.Price)
}

func main() {
	provider := &StockDataProvider{}
	app := &Application{}
	adapter := &XMLToJSONAdapter{app: app}
	client := &Client{}

	// Get XML data from provider
	xmlData := provider.GetStockData()
	fmt.Println("Original XML data:")
	fmt.Println(xmlData)

	// Process XML data (old way)
	xmlStockData, err := app.ProcessStockData(xmlData)
	if err != nil {
		fmt.Println("Error processing XML data:", err)
		return
	}

	// Client uses the XML data
	client.UseXMLData(xmlStockData)

	// Convert XML to JSON using the adapter
	jsonStockData, err := adapter.ConvertToJSON(xmlData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Client uses the JSON data
	client.UseJSONData(jsonStockData)

	// Demonstrate JSON output
	jsonOutput, _ := json.MarshalIndent(jsonStockData, "", "  ")
	fmt.Println("\nJSON output:")
	fmt.Println(string(jsonOutput))
}
