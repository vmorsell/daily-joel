// Daily Joel #7
// Multiple results.
package main

import "fmt"

func concatenate(a, b string) string {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a := "el"
	b := "Jo"

	a, b = swap(a, b)

	fmt.Println(concatenate(a, b))
}
