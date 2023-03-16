package main

import "fmt"

func main() {
	i := 0
	for i < 5 {
		fmt.Printf("Nilai i = %d\n", i)
		i++

		if i == 5 {
			j := 0
			for j <= 10 {
				if j == 5 {
					word := "САШАРВО"
					for idx, val := range word {
						fmt.Printf("character %U '%c' starts at byte position %d\n", val, val, idx)
					}

					j++
					continue
				}

				fmt.Printf("Nilai j = %d\n", j)
				j++
			}
		}
	}
}
