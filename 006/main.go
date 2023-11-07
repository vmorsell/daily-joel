// Daily Joel #6
// Multiple parameters of the same type.
package main

import (
	"fmt"
)

func concatenate(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(concatenate("Jo", "el"))
}
