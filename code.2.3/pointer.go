package main

import "fmt"

var (
	i  int
	pi *int
)

func main() {
	// pi = &i
	// fmt.Scanf("%d", pi)
	// fmt.Println(i, *pi, pi)

	/***** SLICE *****/
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice[2] = 23
	fmt.Println(slice)
}
