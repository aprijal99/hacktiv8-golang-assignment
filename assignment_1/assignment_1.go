package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j)

	fmt.Println() // break

	fmt.Printf("%b\n", i)
	fmt.Println("\u042F")
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Println("\x66")
	fmt.Println("\x46")
	fmt.Printf("%U\n", 'Я')

	fmt.Println() // break

	fmt.Printf("%f\n", 123.456000)
	fmt.Printf("%E\n", 123.456000)
}
