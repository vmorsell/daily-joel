package main

import (
	"fmt"
)

func main() {
	letters := map[bool]string{
		true:  "Joel",
		false: "Viktor",
	}
	for king, v := range letters {
		if king {
			fmt.Println(v)
		}
	}
}
