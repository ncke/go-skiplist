package main

import (
	"SkipList/Skiplist"
	"fmt"
)

func show(s *string) {
	if s == nil {
		fmt.Println("not found")
	} else {
		fmt.Println(*s)
	}
}

func main() {
	fmt.Println("hello")

	var skip = skiplist.Make[int, string](6)

	skip.Insert(5, "world")
	var res = skip.Find(5)
	show(res)

	skip.Insert(2, "cruel")
	res = skip.Find(2)
	show(res)

	res = skip.Find(99)
	show(res)

}
