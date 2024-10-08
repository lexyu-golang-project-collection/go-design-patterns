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

// Client uses the DataConverterInterface
type Client struct {
	converter DataConverterInterface
}

func (c *Client) ProcessData(data []byte) ([]byte, error) {
	return c.converter.Convert(data)
}

// DataConverterInterface defines the interface expected by the client
type DataConverterInterface interface {
	Convert(data []byte) ([]byte, error)
}

// XMLToJSONAdapter adapts the XMLToJSONService to the DataConverterInterface
type XMLToJSONAdapter struct {
	service *XMLToJSONService // adaptee
}

func (a *XMLToJSONAdapter) Convert(data []byte) ([]byte, error) {
	return a.service.ConvertXMLToJSON(data)
}

// XMLToJSONService is our "incompatible" service that we need to adapt
type XMLToJSONService struct{}

func (s *XMLToJSONService) ConvertXMLToJSON(xmlData []byte) ([]byte, error) {
	var xmlStock XMLData
	err := xml.Unmarshal(xmlData, &xmlStock)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}

	jsonStock := JSONData{
		Symbol: xmlStock.Symbol,
		Price:  xmlStock.Price,
	}

	return json.Marshal(jsonStock)
}

func main() {
	// Sample XML data
	xmlString := `
		<stock>
			<symbol>AAPL</symbol>
			<price>150.25</price>
		</stock>
	`

	// Create the service
	service := &XMLToJSONService{}

	// Create the adapter
	adapter := &XMLToJSONAdapter{
		service: service,
	}

	// Create the client with the adapter
	client := &Client{
		converter: adapter,
	}

	// Process the data
	jsonData, err := client.ProcessData([]byte(xmlString))
	if err != nil {
		fmt.Println("Error processing data:", err)
		return
	}

	fmt.Println("Original XML data:")
	fmt.Println(xmlString)
	fmt.Println("Converted JSON data:")
	fmt.Println(string(jsonData))
}
