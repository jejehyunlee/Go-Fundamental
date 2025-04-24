package main

import (
	"fmt"
	"strings"
)

func main() {

	dataString := "Jefri"

	trim := strings.Trim(dataString, "J")
	fmt.Println("Ini Trim :", trim) //Hilangkan Kata

	toLower := strings.ToLower(dataString)
	fmt.Println("Ini ToLower :", toLower) //to Lowercase

	toUpper := strings.ToUpper(dataString)
	fmt.Println("Ini ToUpper : ", toUpper) //to Uppercase

	split := strings.Split(dataString, "f")
	fmt.Println("Ini Split : ", split) //split

	repeat := strings.Repeat(dataString, 3)
	fmt.Println("Ini Repeat : ", repeat) //repeat

	contains := strings.Contains(dataString, "f")
	fmt.Println("Ini Contains : ", contains) //contains

	containsRune := strings.ContainsRune(dataString, 'f')
	fmt.Println("Ini ContainsRune : ", containsRune) //containsRune

	index := strings.Index(dataString, "f")
	fmt.Println("Ini Index : ", index) //index

	indexRune := strings.IndexRune(dataString, 'f')
	fmt.Println("Ini IndexRune : ", indexRune) //indexRune

	count := strings.Count(dataString, "f")
	fmt.Println("Ini Count : ", count) //count

	replace := strings.Replace(dataString, "f", "a", 2)
	fmt.Println("ini replace : ", replace) //replace

	replaceAll := strings.ReplaceAll(dataString, "f", "a")
	fmt.Println("ini replaceAll : ", replaceAll) //replaceAll

}
