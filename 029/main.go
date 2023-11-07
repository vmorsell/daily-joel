// Daily Joel #29
// Methods on non-struct types.
package main

import "fmt"

type Name string

func (n Name) Println() {
	fmt.Println(n)
}

func main() {
	var n Name = "Joel"
	n.Println()
}
