package main

import "fmt"

func main() {

	x := 7

	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Between 2 to 6")
	}

}
