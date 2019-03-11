package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// builtin delete removes key/value pair from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// check if key present in map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize map in same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
