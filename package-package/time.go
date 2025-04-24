package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Before(now))
	fmt.Println(now.After(now))

	layout := time.RFC3339

	parse, _ := time.Parse(layout, "2022-11-11T11:11:11Z")
	fmt.Println(parse)
	fmt.Println(now.Format(layout))

}
