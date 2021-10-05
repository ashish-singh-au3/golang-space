//bool pre declared operators

package main

import "fmt"

func main() {

	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("Equal")
	}
	if 2 != 2 {
		fmt.Println("Not equal")
	}

	//with declaration
	x := 42
	if x == 2 {
		fmt.Println("Values are euqal")

	}
	fmt.Println("Values of  x differ")
}

//With declaration
