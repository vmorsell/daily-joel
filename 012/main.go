// Daily Joel #12
// Slices and range form for loops.

package main

import (
	"fmt"
)

func main() {
	letters := []string{"J", "o", "e", "l"}
	for _, l := range letters {
		fmt.Print(l)
	}
	fmt.Print("\n")
}
