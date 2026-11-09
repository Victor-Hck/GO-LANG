package main

import "fmt"

func main() {
	// Declaracao explicita = voce informa o tipo da variavel
	var variavel string = "Variavel 1"
	var valordepois string
	valordepois = "28"
	var (
		variavel3 string
		variavel4 string
	)

	fmt.Println(variavel)
	fmt.Println(valordepois)
	fmt.Println(variavel3, variavel4)
	
	// inferencia de tipo = o proprio go identifica o tipo com base no valor atribuido
	variavel2 := "Variavel 2"
	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel2)
	fmt.Println(variavel5, variavel6)

	// para declararmos constantes basta trocar o var pelo const ! da pra criar toda a estrutura com o const tbm

	// Trocando ordem de valor de uma variavel, geralmente precisa de uma funcao auxiliar em outras linguagem para isso.
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}