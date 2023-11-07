// Daily Joel #19
//
// Passing structs to functions.
package main

import "fmt"

type Person struct {
	name string
}

func main() {
	p := Person{}
	setName("Joel", &p)
	setNameWithoutEffect("Someone else", p)
	fmt.Println(p.name)
}

func setName(name string, p *Person) {
	p.name = name
}

func setNameWithoutEffect(name string, p Person) {
	p.name = name
}
