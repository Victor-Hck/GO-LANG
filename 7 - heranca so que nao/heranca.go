package main

import "fmt"

// No struct voce pode declarar os valores e nao usar ! agora variaveis voce e obrigado a utilizar.
type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(p1)

	p2 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(p2.altura)
}