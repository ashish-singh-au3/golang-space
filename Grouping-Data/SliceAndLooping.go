package main

import "fmt"

func main() {
	x := []int{1, 5, 6, 7, 8} //declaration

	fmt.Println(x)
	fmt.Println(len(x))

	for i, v := range x { // read as for index, value in the range of x
		fmt.Println(i, v)
	}
}
