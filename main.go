package main

import (
	"SkipList/Skiplist"
	"fmt"
)

func main() {
	fmt.Println("hello")

	var skip = skiplist.Make[string, string](6)
	skip.Insert("hello", "world")
}
