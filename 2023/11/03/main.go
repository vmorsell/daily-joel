// Daily Joel 2023-11-03
//
// Maps.
package main

import (
	"fmt"
	"log"
)

func main() {
	names := map[int]string{
		0: "Joel",
		1: "Someone else",
	}
	n, ok := names[0]
	if !ok {
		log.Fatalf("no element with key 0")
	}
	fmt.Println(n)
}
