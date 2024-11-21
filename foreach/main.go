package main

import "fmt"

func main()  {
	var coding_languages = []string{"Python", "Javascript", "Golang"}
	var tech string = "Golangs"
	var dict = map[string]string {
		"Paris": "France",
		"New York": "USA",
		"Londan": "UK",
	}
	
	for _, value := range(coding_languages){
		fmt.Println(value)
	}

	for _, value := range(dict) {
		fmt.Println(value)
	}

	for _, value:= range(tech) {
		fmt.Printf("%c", value)
	}

	for key, value := range(dict) {
		fmt.Println(key, value)
	}

}