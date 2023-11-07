// Daily Joel #14
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
