//Iota is special pre declared identifier
//it automatically incrementss the value
package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e //iota will start incrementing values automatically
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)

	fmt.Println(d)

	fmt.Println(e)

	fmt.Println(f)

}
