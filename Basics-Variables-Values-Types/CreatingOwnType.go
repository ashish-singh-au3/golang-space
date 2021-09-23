package main

import "fmt"

type ashish int //created own type named ashish

var b ashish

func main() {
	b = 20

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
