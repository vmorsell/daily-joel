// Daily Joel #27
// Maps and zero values.
package main

import "fmt"

func main() {
	names := map[int]string{
		2: "Joel",
		4: "\n",
	}

	for k := 0; k < 10; k++ {
		fmt.Print(names[k])
	}
}
