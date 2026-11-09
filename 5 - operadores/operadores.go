package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, divisao)

	// go e fortimente tipado ! voce nao consegue somar nem fazer nada com valores diferentes
	var numero1 int16 = 10 
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMETICOS

	// ATRIBUICAO
	var variavel1 string = "atribuido"
	variavel2 := "string"
	fmt.Println(variavel1, variavel2)
	// FIM OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	// FIM OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // and
	fmt.Println(verdadeiro || falso) // or
	fmt.Println(!verdadeiro) // negacao
	fmt.Println(!falso) // negacao
	// FIM OPERADORES LOGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 10
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNARIOS

	// OPERADOR TERNARIO -> NAO EXISTE NO GO
	// POIS O GO FOCA EM SO 1 MANEIRA DE FAZER CADA COISA
	// texto := numero > 5 ? "maior" : "menor"
	// fmt.Sprintln(texto)
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto) // maneira de fazer um operador ternario


}
