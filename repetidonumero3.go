package main

import (
	"errors"
	"fmt"
)

func main() {
	// numeros inteiros
	// todos os types de numeros 
	var numero int = 10000000 // se voce declarar apenas o int ele usa de base o seu computador 64 ou 32 bits
	// int8, int16, int32, int64 // diferença deles esta nos bits que eles ocupam
	fmt.Println(numero)

	// uint // unsygned int = um int sem sinal
	var numero2 uint32 = -1000 // nao pode usar o valor negativo
	fmt.Println(numero2) 

	// todos os tipos possuiem suas alias
	// int32 = rune obs: geralmente utilizado com numeros que representam caracteres utilizados pela tabela ASK
	var numero3 rune = 123 // mesma coisa de um int32
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// numeros reais float32, float64
	var numeroReal float32 = 67.00 // precisa ser sempre o ponto. se for virgula nao vai funcionar
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 67.00000000000000000 // precisa ser sempre o ponto. se for virgula nao vai funcionar
	fmt.Println(numeroReal2)

	// existe uma diferenca entre "" e '' no go
	// char = retorna o numero que o caracter esta na tabela ask
	char := 'B'
	fmt.Println(char)

	// na linguagem go existe om conceito que se chama VALOR ZERO
	// basicamente toda variavel que nao possui valor quando voce cria ela retorna o valor zero por padrao seja ums tring como vazio e um inteiro com 0
	// todo tipo de dado tem seu valor inicial que e o VALOR ZERO
	// FIM STRING
	var texto string
	fmt.Println(texto)

	// tipos boloanos
	var booleano1 bool = false
	var booleano2 bool = true

	// tipo error
	// erro = nome variavel error = tipo errors = pacote
	var erro error = errors.New("error interno")
	fmt.Println(erro) // retorna <nil> nao de nulo mais em de zero


}