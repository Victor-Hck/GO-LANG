package main

import "fmt"

func main() {
	// if initi ! quando você consegue iniciar uma variavel no if
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 { // só se usa () para utilizar ordem de precedencia na hora de fazer alguma conta
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual a 15")
	}

	if numero > 15 { fmt.Println("Maior que 15") }

	// Quando se cria uma variavel usando o if init você limita ela ao escopo dela. Não sendo possivel chamar ela globalmente.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}else {
		fmt.Println("Numero é menor que zero")
	}

}

