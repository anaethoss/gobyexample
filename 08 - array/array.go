package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var c = [5]int{5, 4, 3, 2, 1}
	fmt.Println("C:", c)

	var twoD [2][5]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

	var fourD [2][3][4][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 3; l++ {
					fourD[i][j][k][l] = i + j + k + l
				}
			}
		}
	}

	fmt.Println("FourD:", fourD)

}
