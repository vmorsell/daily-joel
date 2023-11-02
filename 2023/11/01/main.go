// 2023-11-01
//
// Appending to slices.
package main

import "fmt"

func main() {
	name := "Joel"

	runes := []rune{}
	for _, r := range name {
		runes = append(runes, r)
	}

	fmt.Println(string(runes))
}
