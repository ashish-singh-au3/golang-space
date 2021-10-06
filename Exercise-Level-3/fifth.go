/* Create a program that shows the “if statement” in action. */

package main

import "fmt"

func main() {
	x := 10
	if x == 2 {
		fmt.Println("The values don't match")
	}
	fmt.Println("The correct value is", x)
}
