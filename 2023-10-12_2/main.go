package main

import (
	"fmt"

	"github.com/vmorsell/joel/2023-10-12/names"
)

func main() {
	var name string

	fmt.Println("what is your name?")

	fmt.Scan(&name)

	if name == names.GitKenobi || name == names.Vmorsell {
		fmt.Printf("Welcome, %s \n",
			name)
	} else {
		fmt.Println("You don't live here, leave")
	}
}
