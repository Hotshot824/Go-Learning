package main

import "fmt"

func main() {
	i, j := 1, 1
	for i <= 9 {
		j = 1
		for j <= 9 {
			fmt.Print(i * j)
			fmt.Print(" ")
			j++
		}
		fmt.Print("\n")
		i++
	}
}
