// Daily Joel 2023-10-20
//
// Arrays and basic for loops.
package main

import (
	"fmt"
)

func main() {
	letters := [4]string{"J", "o", "e", "l"}
	for i := 0; i < len(letters); i++ {
		fmt.Print(letters[i])
	}
	fmt.Print("\n")
}
