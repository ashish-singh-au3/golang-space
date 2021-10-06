/* Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
 */

package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James" {
		fmt.Println("this statement is wrong")
	} else if x == "Bond" {
		fmt.Println("James is missing")
	} else {
		fmt.Println("the correct value is", x)
	}
}
