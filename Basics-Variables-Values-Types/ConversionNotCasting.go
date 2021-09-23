package main

import "fmt"

var a int

type ashish int

var b ashish

func main() {
	a = 10
	b = 20

	fmt.Println(a)        //10
	fmt.Printf("%T\n", a) //int

	fmt.Println(b)        //20
	fmt.Printf("%T\n", b) // int

	a = int(b)     //type conversion since we already have declare var a as int in line no 5
	fmt.Println(a) //10
	fmt.Printf("%T\n", a)

}
