package main

import "fmt"

func main() {
	fmt.Println(CountLetters("Hello World!"))
	fmt.Println(CountLetters("123 456"))
	fmt.Println(CountLetters("Golang1.21"))
	fmt.Println(CountLetters("école"))
}

func CountLetters(s string) int {
	compteur := 0
	for i := 0; i != len(s); i++ {
		if s[i] > 65 && s[i] < 122 {
			compteur += 1
		}
	}
	return compteur
}
