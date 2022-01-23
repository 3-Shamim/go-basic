package main

import "fmt"

func main() {

	//var a [5]int

	//a := [5]int{1, 2, 5, 6, 8}

	a := []int{1, 2, 5, 6, 8}

	a[2] = 2
	a = append(a, 12)

	fmt.Println(a)

}
