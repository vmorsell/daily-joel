package main

import (
	"fmt"
	"slices"
)

var letters = map[int]string{
	1: "J",
	2: "o",
	3: "e",
	4: "l",
	5: "\n",
}

func main() {
	//quick assigning two indexes for the _horrible_ nested for loops incoming
	is := 0
	ik := 0

	//a for loop with the intended function to print each value by printing the lowest key value for each iteration
	//	¯\_(ツ)_/¯
	for {
		if len(letters) == 0 {
			break
		}

		//I need the key value pairs as an array to use the slice.Min func.
		//I also need this array to be recalculated for each iteration of the surrounding for loop.

		keys := make([]int, len(letters))
		for key := range letters {
			keys[is] = key
			is++
		}
		//assigning variable i to the lowest element in keys slice
		ik = slices.Min(keys)

		//printing the value for the lowest key
		fmt.Print(letters[ik])

		//deleting key value pair that has been printed
		delete(letters, ik)

		//resetting our keys slice by giving it a nil value
		keys = nil
		//resetting the indexes
		is = 0
		ik = 0
		//break if letters is empty

	}
}
