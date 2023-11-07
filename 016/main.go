// Daily Joel #16
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
