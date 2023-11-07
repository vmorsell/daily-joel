// Daily Joel #18
//
// Structs.
package main

import "fmt"

type Person struct {
	name *string
}

func main() {
	n := "Joel"
	p := Person{
		name: &n,
	}
	fmt.Println(*p.name)
}
