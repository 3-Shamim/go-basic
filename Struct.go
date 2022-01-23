package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p := person{name: "Roni", age: 14}
	fmt.Println(p)
	fmt.Println(p.name)
	p.age = 20
	fmt.Println(p.age)

}
