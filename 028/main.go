// Daily Joel #28
// Functions returning functions.
package main

import "fmt"

func main() {
	w := writer()
	w("Joel")
}

func writer() func(s string) {
	return func(s string) {
		fmt.Println(s)
	}
}
