package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should print false")

	case (2 == 4):
		fmt.Println("This shouldn't print 2")
	case (3 == 3):
		fmt.Println("Prints")
		fallthrough
	case (4 != 4):
		fmt.Println("This is wrong")
	case (4 == 4):
		fmt.Println("this is true")

	}

}
