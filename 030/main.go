// Daily Joel #30
// Interfaces.
package main

import "fmt"

type Namer interface {
	Name() string
}

type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func printName(n Namer) {
	fmt.Println(n.Name())
}

func main() {
	j := Person{"Joel"}
	printName(j)
}
