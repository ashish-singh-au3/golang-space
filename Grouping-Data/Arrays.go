package main

import "fmt"

func main() {
	var x [5]int //initialised var of int type with 5 val
	fmt.Println(x)

	x[3] = 42 //inserting 42 at index 3
	fmt.Println(x)

}
