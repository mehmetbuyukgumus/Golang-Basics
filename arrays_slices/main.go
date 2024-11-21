package main

import "fmt"

func main() {
	// Fixed
	var names [3]string
	names[0] = "Mehmet"
	names[1] = "İlker"
	names[2] = "Furkan"

	var last_names = [3]string{"Büyükgümüş", "Uslu", "Arabacı"}
	names[1] = "Alper"
	
	fmt.Println(names)
	fmt.Println(last_names)

	// None-fixed
	var team = []string{"Icardi", "Osimhen", "Mertens", "Yunus", "Sara"}
	// Adding element
	team = append(team, "Barış")
	fmt.Println(team)

	var numbers = []int{1,2,3,4}
	fmt.Println(numbers)

	//Slices
	fmt.Println(team[0:2])
}
