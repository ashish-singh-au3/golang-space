package main

import "fmt"

func main() {
	x := 42
	if x == 5 {
		fmt.Println("Values is 5")
	} else if x == 41 {
		fmt.Println("Val is 41")
	} else {
		fmt.Println("value isn't 42")
	}
}
