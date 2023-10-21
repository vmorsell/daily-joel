// Daily Joel 2023-10-13
//
// Sum of strings.
package main

import (
	"fmt"
)

func appendEl(s string) string {
	return s + "el"
}

func main() {
	fmt.Println(appendEl("Jo"))
}
