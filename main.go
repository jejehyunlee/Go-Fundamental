package main

import (
	"fmt"
	"strings"
)

func twoSumNew(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func namaHewan() string {

	data := []string{"babi", "Burung", "Kucing"}

	Data := []bool{true, false, false}
	size := Data[2]
	fmt.Println(size)

	return strings.Join(data, ",")
}

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	target := 8

	data := [...]int{1, 2, 3, 5}
	fmt.Println(data[0]) // Akses elemen pertama

	nums := []int{1, 2, 3, 5}
	target = 8

	hasil := twoSumNew(nums, target)

	fmt.Println(hasil)

	fmt.Println(namaHewan())

}
