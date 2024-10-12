package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// XMLData represents the structure of our XML data
type XMLData struct {
	XMLName xml.Name `xml:"stock"`
	Symbol  string   `xml:"symbol"`
	Price   float64  `xml:"price"`
}

// JSONData represents the structure of our JSON data
type JSONData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

// Client uses both XML and JSON data
type Client struct {
	xmlService  *XMLService
	jsonAdapter *XMLToJSONAdapter
}

func (c *Client) ProcessXMLData(data []byte) {
	xmlStock, err := c.xmlService.ParseXML(data)
	if err != nil {
		fmt.Println("Error processing XML:", err)
		return
	}
	fmt.Printf("Processed XML Data: Symbol=%s, Price=%.2f\n", xmlStock.Symbol, xmlStock.Price)
}

func (c *Client) ProcessJSONData(data []byte) {
	jsonStock, err := c.jsonAdapter.ConvertToJSON(data)
	if err != nil {
		fmt.Println("Error processing JSON:", err)
		return
	}
	fmt.Printf("Processed JSON Data: Symbol=%s, Price=%.2f\n", jsonStock.Symbol, jsonStock.Price)
}

// XMLService handles XML data
type XMLService struct{}

func (s *XMLService) ParseXML(data []byte) (*XMLData, error) {
	var xmlStock XMLData
	err := xml.Unmarshal(data, &xmlStock)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}
	return &xmlStock, nil
}

// XMLToJSONAdapter adapts XML data to JSON
type XMLToJSONAdapter struct {
	xmlService *XMLService
}

func (a *XMLToJSONAdapter) ConvertToJSON(xmlData []byte) (*JSONData, error) {
	xmlStock, err := a.xmlService.ParseXML(xmlData)
	if err != nil {
		return nil, err
	}

	jsonStock := &JSONData{
		Symbol: xmlStock.Symbol,
		Price:  xmlStock.Price,
	}

	return jsonStock, nil
}

func main() {
	// Sample XML data
	xmlString := `
		<stock>
			<symbol>AAPL</symbol>
			<price>150.25</price>
		</stock>
	`

	xmlService := &XMLService{}
	jsonAdapter := &XMLToJSONAdapter{xmlService: xmlService}
	client := &Client{
		xmlService:  xmlService,
		jsonAdapter: jsonAdapter,
	}

	fmt.Println("Original XML data:")
	fmt.Println(xmlString)

	// Client processes XML data
	fmt.Println("\nProcessing XML:")
	client.ProcessXMLData([]byte(xmlString))

	// Client processes JSON data (converted from XML)
	fmt.Println("\nProcessing JSON (converted from XML):")
	client.ProcessJSONData([]byte(xmlString))

	// Demonstrate JSON output
	jsonStock, _ := jsonAdapter.ConvertToJSON([]byte(xmlString))
	jsonOutput, _ := json.MarshalIndent(jsonStock, "", "  ")
	fmt.Println("\nJSON output:")
	fmt.Println(string(jsonOutput))
}
