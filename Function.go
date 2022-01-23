package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(10, 20))
	result := sum(10, 20)
	fmt.Println(result)

	res, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil

}
