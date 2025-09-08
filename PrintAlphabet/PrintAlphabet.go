package main

import "fmt"

func main() {
	PrintAlphabet()
}

func PrintAlphabet() {
	for i := a; i < z; i++ {
		fmt.Println("%c", i)
	}
	fmt.Println("\n")
}
