package main

import "fmt"

func main() {

	var condition = true

	defer cleanUp()
	if condition {
		panic("An error occured")
	}

}

func cleanUp() {
	fmt.Println("File has been cleaned")
}
