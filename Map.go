package main

import "fmt"

func main() {

	combo := make(map[string]int)

	combo["a"] = 2
	combo["b"] = 5
	combo["g"] = 7

	fmt.Println(combo["b"])

	delete(combo, "b")

	fmt.Println(combo)

}
