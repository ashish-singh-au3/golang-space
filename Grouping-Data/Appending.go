package main

import "fmt"

func main() {
	x := []int{5, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 10, 20, 30)
	fmt.Println(x) //5,6,7,8,10,20,30

	y := []int{110, 121, 130}
	x = append(x, y...) //taking out the values of y and appending it in x
	fmt.Println(x)
}
