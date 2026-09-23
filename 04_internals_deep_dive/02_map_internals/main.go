package main

import (
	"fmt"
)

func main() {
	langs := map[string]int{
		"Go":   2009,
		"Java": 1995,
	}
	fmt.Println(langs)
	modifyMap(langs)
	fmt.Println(langs)
}

func modifyMap(m map[string]int) {
	m["C++"] = 1983
	delete(m, "Java")
}
