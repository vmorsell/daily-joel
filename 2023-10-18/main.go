package main

import "fmt"

func concatenate(a, b string) string {
	return a + b
}

func main() {
	a := "el"
	b := "Jo"

	a, b = b, a

	fmt.Println(concatenate(a, b))
}
