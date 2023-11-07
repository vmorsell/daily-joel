// Daily Joel #20
//
// Methods.
package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p Person) setNameWithoutEffect(name string) {
	p.name = name
}

func main() {
	p := Person{}
	p.setName("Joel")
	p.setNameWithoutEffect("Someone else")
	fmt.Println(p.name)
}
