// Daily Joel #26
// Mutating maps.
package main

import (
	"fmt"
	"log"
)

func main() {
	names := make(map[int]string)
	names[0] = "Joel"
	names[1] = "Someone else"
	n, ok := names[0]
	if !ok {
		log.Fatalf("key not found")
	}
	fmt.Println(n)
}
