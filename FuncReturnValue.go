package main

import (
	"fmt"
)

func duaAngka(data []int, target int) []int {

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {

	data := []int{1, 2, 3, 5}
	target := 8

	//hasil := duaAngka(data, target)

	fmt.Println(duaAngka(data, target))

}
