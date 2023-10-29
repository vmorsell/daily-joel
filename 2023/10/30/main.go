// Daily Joel 2023-10-31
//
// Slices and zero values.
package main

import "fmt"

func main() {
	letters := make([]string, 100)
	letters[2] = "J"
	letters[13] = "o"
	letters[37] = "e"
	letters[71] = "l"
	letters[95] = "\n"

	for _, l := range letters {
		fmt.Print(l)
	}
}
