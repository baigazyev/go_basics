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

func main() {
	// Create a Product instance
	product := Product{Name: "Laptop", Price: 999.99, Quantity: 10}

	// Convert Product to JSON
	jsonString, err := ProductToJSON(product)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
	} else {
		fmt.Println("Product in JSON format:", jsonString)
	}
}

