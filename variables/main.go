package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Method 1
	var productName1 string = "PC"
	var quantitiy1 int = 12
	var discount1 float32 = 0.37
	var isInStock1 bool = true

	fmt.Println(productName1, quantitiy1, discount1, isInStock1)
	fmt.Println(reflect.TypeOf(productName1), reflect.TypeOf(quantitiy1), reflect.TypeOf(discount1), reflect.TypeOf(isInStock1))

	// Method 2
	var productName2 string = "PC"
	var quantitiy2 int = 12
	var discount2 float32 = 0.37
	var isInStock2 bool = true

	fmt.Println(productName2, quantitiy2, discount2, isInStock2)
	fmt.Println(reflect.TypeOf(productName2), reflect.TypeOf(quantitiy2), reflect.TypeOf(discount2), reflect.TypeOf(isInStock2))

	// Method 3
	productName3 := "PC"
	quantitiy3 := 12
	discount3 := 0.37
	isInStock3 := true

	fmt.Println(productName3, quantitiy3, discount3, isInStock3)
	fmt.Println(reflect.TypeOf(productName3), reflect.TypeOf(quantitiy3), reflect.TypeOf(discount3), reflect.TypeOf(isInStock3))

	// Variable Formatting
	fmt.Printf("Product Name: %s,Discount: %f\nQuantity: %d,\nDiscount: %f,\nIs in stock: %t", productName3,discount3, quantitiy3, discount3, isInStock3)
}
