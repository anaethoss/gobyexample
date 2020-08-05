package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "hash this string with sha1"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(h)
	fmt.Printf("%x\n", bs)
}
