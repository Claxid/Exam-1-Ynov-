package main

import "fmt"

func main() {
	fmt.Println(PrintPGCD(14, 21))
}

func PrintPGCD(a int, b int) int {
	pgcd_f := 1
	if a < b {
		for i := b; i > 0; i-- {
			if a%i == 0  {
			 	if b%i == 0 {
					pgcd_f = i
			}
		}
	} else {
		for i := a; i > 0; i-- {
			if a%i == 0 {
				if b%i == 0 {
					pgcd_f = i

			}
		}
	

	}

}
}
return pgcd_f
}