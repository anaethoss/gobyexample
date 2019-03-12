package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}

	sum := 0

	// _ for ignore index of iterator
	for _, num := range nums {
		sum += num
		fmt.Println(num)
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for K := range kvs {
		fmt.Println("Key:", K)
	}

	for _, Val := range kvs {
		fmt.Println("Value:", Val)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
