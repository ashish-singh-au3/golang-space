/* Create TYPED and UNTYPED constants. Print the values of the constants. */

package main

import "fmt"

const (
	a     = 40 //untyped constant
	b int = 30 //typed const
)

func main() {
	fmt.Println(a, b)
}
