// Daily Joel #13
//
// Switch.
package main

import (
	"fmt"
)

func name(initial string) string {
	switch initial {
	case "j":
		return "Joel"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(name("j"))
}
