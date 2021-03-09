package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conversor - Binário para Decimal")
	fmt.Print("\nDigite o número a ser convertido: ")

	var entrada string
	fmt.Scan(&entrada)

	saida := conversor(entrada)
	fmt.Printf("\nResultado:\nBinário: %s\nDecimal: %d\n", entrada, saida)
}
