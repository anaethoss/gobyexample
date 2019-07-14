package main

import "fmt"

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {

	printInt := intSeq()

	fmt.Println(printInt()) // 1
	fmt.Println(printInt()) // 2
	fmt.Println(printInt()) // 3
	fmt.Println(printInt()) // 4
	fmt.Println(printInt()) // 5

	newInt := intSeq()

	fmt.Println(newInt()) // 1
}
