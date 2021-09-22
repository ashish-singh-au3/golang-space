package main

import "fmt"

var x = 20

//Declared the variable with identifier y
// is of Type String
//Assigned the value "This is to check the type and the type is "
var y string = "This is to check the type and the type is "

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
