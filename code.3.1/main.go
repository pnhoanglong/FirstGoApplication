package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d", &x)

	if x > 18 {
		fmt.Println("Congratulations!")
		return
	}

	fmt.Println("Sorry, you are too young")
}
