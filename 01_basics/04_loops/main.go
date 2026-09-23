package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 1; i <= 50; i++ {
		sum += i
	}

	fmt.Println(sum)

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println("Go!")

}
