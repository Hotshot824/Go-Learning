package table99

import "fmt"

func Table99() {
	i, j := 1, 1
	for i <= 9 {
		j = 1
		for j <= 9 {
			fmt.Print(i * j)
			fmt.Print("\t")
			j++
		}
		fmt.Print("\n")
		i++
	}
}
