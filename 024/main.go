// Daily Joel #24
// Breaking loops.
package main

import "fmt"

func main() {
	letters := []string{"J", "o", "e", "l", "\n", "x", "y", "z"}
	breakpoint := 5
	for i, l := range letters {
		if i == breakpoint {
			break
		}
		fmt.Print(l)
	}
}
