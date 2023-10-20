// Daily Joel 2023-10-19
//
// Variable declaration and initialization, variable scopes,
// and variadic parameters.
package main

import (
	"fmt"
)

func concatenate(s ...string) string {
	out := ""
	for _, ss := range s {
		out += ss
	}
	return out
}

var a, b string = "N", "o"
var c string = "e"
var d string

func printJoel() {
	a := "J"
	d = "l"
	fmt.Println(concatenate(a, b, c, d))
}

func main() {
	printJoel()

	// What's the current values of the variables?
}
