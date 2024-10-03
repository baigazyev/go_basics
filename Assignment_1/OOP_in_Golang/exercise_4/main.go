package main

// Import the necessary packages
import (
	"encoding/json"
	"fmt"
)

// Define the Product struct
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func ProductToJSON(p Product) (string, error) {
	// Marshal the struct into JSON format
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JSONToProduct(jsonString string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func main() {
	// Create a Product instance
	product1 := Product{Name: "Laptop", Price: 999.99, Quantity: 10}
	jsonString1 := `{"name":"Laptop","price":999.99,"quantity":10}`

	// Convert Product to JSON
	jsonString2, err := ProductToJSON(product1)

	// Convert JSON back to Product
	product2, err := JSONToProduct(jsonString1)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product in JSON format:", jsonString2)
		fmt.Println("JSON in product format:", product2)
	}
}
