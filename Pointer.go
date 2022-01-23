package main

import "fmt"

func main() {

	i := 7
	fmt.Println(i)
	fmt.Println(&i)
	inc(i)
	fmt.Println(i)
	inc1(&i)
	fmt.Println(i)

}

func inc(x int) {
	x++
}

func inc1(x *int) {
	*x++
}
