package main

import (
	"fmt"
	"math"
)

func conversor(e string) int {
	e = reverse(e)
	contador := 0

loop:
	for i := 0; i < len(e); i++ {
		elemento := string(e[i])

		switch elemento {
		case "1":
			contador += int(math.Pow(2, float64(i)))
		case "0":
			continue
		default:
			fmt.Println("\nErro!\nO número a ser convertido deve ser um número binário\n ")
			break loop
		}
	}

	return contador
}
