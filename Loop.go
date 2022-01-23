package main

import "fmt"

func main() {

	fmt.Println("====================== 1")

	// 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("====================== 2")

	//2
	j := 0

	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("======================= 3")

	// 3
	arr := []string{"a", "b", "c", "d"}

	for index, value := range arr {
		fmt.Printf("index: %d ==== value: %s", index, value)
		fmt.Println()
	}

	fmt.Println("======================= 4")

	// 3
	m := make(map[string]string)
	m["a"] = "aa"
	m["b"] = "bb"

	for key, value := range m {
		fmt.Printf("key: %s ==== value: %s", key, value)
		fmt.Println()
	}

}
