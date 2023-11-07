// Daily Joel #15
//
// Stacking defers.
package main

import "fmt"

func main() {
	letters := []string{"\n", "l", "e", "o", "J"}
	for _, l := range letters {
		defer fmt.Print(l)
	}
}
