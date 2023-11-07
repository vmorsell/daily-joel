// Daily Joel #22
// Slices of slices.
package main

import "fmt"

func main() {
	letters := [][]string{
		{"J", "o"},
		{"e"},
		{"l", "\n"},
	}
	for _, l := range letters {
		for _, ll := range l {
			fmt.Print(ll)
		}
	}
}
