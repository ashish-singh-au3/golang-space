package main

import "fmt"

func main() {

	x := 4

	fmt.Printf("%d\t\t%b\n", x, x) // output 2 is decimal number system and 10 is binary numbering system value

	y := x << 1 //Bit shifting process

	fmt.Printf("%d\t\t%b", y, y)
}
