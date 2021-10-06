/* Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive. */

package main

import "fmt"

/* func main() {
	bd := 1997
	for bd <= 2021 {
		fmt.Println(bd)
		bd++
	}
} */

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive. */

func main() {
	bd := 1997
	for {
		if bd > 2021 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
