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

func Table99plus() {
	count := 0
	for true{
		if count > 7{
			break
		}
		for i := 1; i < 10; i++{
			for j := 1 + count; j < 4 + count; j++{
				fmt.Print(j, " * ", i, " = ", j * i, "\t")
			}
			fmt.Print("\n")
		} 
		count += 3
	}
}
