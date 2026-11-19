package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAYS	
	fmt.Println("Arrays e Slices")

	var array1 [5]string // se so printar um array vazio ele retorna zero
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	
	// ARRAYS QUASE NAO SAO USADOS EM GO POIS LIMITA O TAMANHO DAS COISAS
	// SLICE
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Sprintln(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	array2[1] = "Posicao Alterada"
	fmt.Println(slice2)

	// Arrays Internos Entender melhor sobre a função Make
	fmt.Println("----------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}