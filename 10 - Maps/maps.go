package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	usuario := map[string]string { // maps é mais rigido em relação ao tipo do dado
		"Nome": "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario["Nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Antonio",
			"ultimo": "Gomes",
		},
		"curso": {
			"nome": "Redes",
			"campus": "campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") // Deletando uma chave
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}