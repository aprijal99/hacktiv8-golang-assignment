package main

import (
	"fmt"
)

func main() {
	name := "selamat malam"
	calculateLetter(name)
}

func calculateLetter(name string) {
	letterMap := make(map[string]int)

	for _, val := range name {
		fmt.Printf("%c\n", val)

		var valTmp string = string(val)
		letterMap[valTmp]++
	}

	fmt.Println(letterMap)
}
