package main

import (
	"errors"
	"fmt"
)

func main() {
	// Example 1
	result, err := divide(10, 1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	productService := ProductService{}
	err = productService.Add(Product{id: 1, name: "Pc", price: 3000})

	if err != nil {
		fmt.Println(err)
	}

}

// Example 1
func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("payda sıfır olamaz")
	}
	return x / y, nil
}

// Example 2
type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return errors.New("product name can not be blank")
	}
	if product.price <= 0 {
		return errors.New("product's price can't be zero or lower than zero")
	}
	fmt.Println("Product has been added")
	return nil
}
