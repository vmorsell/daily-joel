// Daily Joel 2023-10-25
//
// Pointers.
package main

import (
	"fmt"
)

func main() {
	name := ""

	p := &name
	*p = "Joel"

	fmt.Println(name)
}
