package main

import "fmt"

//Declaration of variable "z"
//Assigning the value = 10
//declaration and assigning = initialization

var z = 10

func main() {
	x := 30 //If we declare short hand outside function then it will throw error
	fmt.Println(x)

	var y = 20 //if we declare this even outside the function this will work fine, below is the example
	fmt.Println(y)

	fmt.Println(z) //We have declared var outside func but it will work fine since we have used var

	foo()

}

func foo() {
	fmt.Println("THis will show the same value of z that is :", z)
}
