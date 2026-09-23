package main

import (
	"fmt"
)

func main() {
	fruits := map[string]int{
		"Apples":  10,
		"Bananas": 5,
		//"Pears":   2,
	}

	fruits["Oranges"] = 20

	if pearsCount, ok := fruits["Pears"]; ok {
		fmt.Printf("Pears count is: %d", pearsCount)
	} else {
		fmt.Println("Товар не найден")
	}
}
