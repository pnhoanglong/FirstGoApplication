package main

import "fmt"

/*
	Global (package level) varibale
*/

// variable
var phoneNumber string = "0934 06 26 75"

// block variable
var (
	name string
	age  int
)

// Constant
const bornYear = 1991

// Constant block
const (
	nationality = "Vietnamese"
	religion    = "None"
)

func main() {
	name = "Long Pham"
	age = 2018 - bornYear
	address := "36/5 Nguyen Van Lac, F19 QBT HCM"
	fmt.Printf("Hello world, I am %s, was born %d, %d years old.\n I am %s, my religion is %s. \n I am living at %s\n", name, bornYear, age, nationality, religion, address)
	fmt.Println(address)
}
