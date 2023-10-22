// Daily Joel 2023-10-23
//
// Defer.
package main

import "fmt"

func main() {
	name := ""

	defer func() {
		fmt.Println(name)
	}()

	name = "Joel"
}
