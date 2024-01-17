package main

import (
	"fmt"
)

var namearr = []string{"J", "o", "e", "l", "\n"}

func concatenate(namearr []string, i int) string {
	if i >= len(namearr) {
		return ""
	} else {
		return namearr[i] + concatenate(namearr, i+1)
	}
}

func main() {
	fmt.Print(concatenate(namearr, 0))
}
