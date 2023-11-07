// Daily Joel #8
// Named return values.
package main

import "fmt"

func concatenate(a, b string) string {
	return a + b
}

func swap(a, b string) (x, y string) {
	x = b
	y = a
	return
}

func main() {
	a := "el"
	b := "Jo"

	a, b = swap(a, b)

	fmt.Println(concatenate(a, b))
}
