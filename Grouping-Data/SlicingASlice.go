package main

import "fmt"

func main() {
	x := []int{5, 7, 2, 3, 1}

	fmt.Println(x)
	fmt.Println(x[2])   // 2
	fmt.Println(x[1:])  // 7,2,3,1
	fmt.Println(x[2:4]) //2,3
}
