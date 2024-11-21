package main

import "fmt"


func main() {
	var customer_1 Customer = Customer{id: 1, name: "MyName", age: 30, adress: Adress{city: "London", disctrict: "Zone 5"}}
	customer_1.ToString()
	customer_1.changeName("HisName")
	customer_1.ToString()
}

type Customer struct {
	id int64
	name string
	age int
	adress Adress
}

type Adress struct {
	city string
	disctrict string
}

func (customer Customer) ToString() {
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}