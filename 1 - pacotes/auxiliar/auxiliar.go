package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
}

// Go não é uma linguagem orientada a objetos
// Existe a questão de letras minusculas e maiusculas para os pacotes !
// Se escrever com letra minuscula o pacote só e visto dentro do arquivo que ele esta
// Se escrever com letra maiuscula o pacote pode ser exportado para outros arquivos
// Comentarios em cima de funções dizem o que ela faz. Famoso docstring
// Entender melhor essa questão de arquivos e pacotes ! exporta pacote etc.S