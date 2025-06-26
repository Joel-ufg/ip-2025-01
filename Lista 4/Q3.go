package main

import "fmt"

func main() {
	const tamanho = 10
	vetor := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Scan(&vetor[i])
	}

	var pares []int
	var impares []int
	somaPares := 0

	for _, valor := range vetor {
		if valor%2 == 0 {
			pares = append(pares, valor)
			somaPares += valor
		} else {
			impares = append(impares, valor)
		}
	}

	fmt.Println("Números pares digitados:", pares)
	fmt.Println("Soma dos números pares digitados:", somaPares)
	fmt.Println("Números ímpares digitados:", impares)
	fmt.Println("Quantidade de números ímpares digitados:", len(impares))
}
