// Daily Joel #5
//
// Multiple parameters.
package main

import (
	"fmt"
)

func concatenate(a string, b string) string {
	return a + b
}

func main() {
	fmt.Println(concatenate("Jo", "el"))
}
