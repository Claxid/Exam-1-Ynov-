package main

import "fmt"

func main() {
	PrintDecimal()
}

func PrintDecimal() {
	for i := 49; i <= 57; i++ {
		fmt.Printf("%c", i)
		fmt.Printf("\n")
	}
}
