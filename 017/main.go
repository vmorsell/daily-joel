// Daily Joel #17
//
// Forever loops.
package main

import (
	"fmt"
)

func main() {
	letters := []string{"J", "o", "e", "l", "\n"}
	for {
		if len(letters) == 0 {
			break
		}
		fmt.Print(letters[0])
		letters = letters[1:]
	}
}
