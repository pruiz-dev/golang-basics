package main

import (
	"fmt"
)

func main() {

	langs := []string{"Go", "Python", "Java"}
	langs = append(langs, "C++")

	for i, v := range langs {
		fmt.Printf("Language #%d: %s\n", i, v)
	}

}
