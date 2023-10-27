// Daily Joel 2023-10-28
//
// Passing structs to functions.
package main

import "fmt"

type Person struct {
	name string
}

func main() {
	j := Person{}
	setName("Joel", &j)
	setNameWithoutEffect("Someone else", j)
	fmt.Println(j.name)
}

func setName(name string, p *Person) {
	p.name = name
}

func setNameWithoutEffect(name string, p Person) {
	p.name = name
}
