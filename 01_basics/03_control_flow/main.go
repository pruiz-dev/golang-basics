package main

import (
	"fmt"
)

func main() {
	score := 4
	message := "Abnormal situation"

	if score < 0 || score > 100 {
		fmt.Println(message)
		return
	}

	if score < 70 {
		message = "Needs improvement."
	} else if score < 100 {
		message = "Good job!"
	} else { // score is equal 100
		message = "Perfect!"
	}

	fmt.Println(message)

}
