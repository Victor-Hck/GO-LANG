package main

import "fmt"

// struct e como se fosse um classe em outras linguagens !

type usuario struct {
	nome string
	idade uint8
	endereco endereco // endereco do tipo endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario // populando struct
	u.nome = "akadi"
	u.idade = 25
	fmt.Println(u) // no print so aparece o valor nao aparece nome valor !

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// outra maneira de fazer a criacao de um struct
	usuario2 := usuario{"Davi", 25, enderecoExemplo}
	fmt.Println(usuario2)

	// 3 forma de fazer
	usuario3 := usuario{nome: "davi"} // usasse quando nao se tem o dado que vai colocar para os outros tipos.
	fmt.Println(usuario3)
}