package main

import "fmt"

func main() {
	PrintDecimal()
}

func PrintDecimal() {
	fmt.Println("0")
	for i := 49; i <= 57; i++ {
		fmt.Printf("%c", i)
		fmt.Printf("\n")
	}
}
