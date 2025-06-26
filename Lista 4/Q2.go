package main

import "fmt"

func main() {

	const tamanho1 = 10
	const tamanho2 = 5

	vetor1 := make([]int, tamanho1)
	vetor2 := make([]int, tamanho2)

	for i := 0; i < tamanho1; i++ {
		fmt.Scan(&vetor1[i])
	}

	for i := 0; i < tamanho2; i++ {

		fmt.Scan(&vetor2[i])
	}

	somaVetor2 := 0
	for _, v := range vetor2 {
		somaVetor2 += v
	}

	var paresComSoma []int
	var imparesComSoma []int

	for _, valor := range vetor1 {
		if valor%2 == 0 {
			paresComSoma = append(paresComSoma, valor+somaVetor2)
		} else {
			imparesComSoma = append(imparesComSoma, valor+somaVetor2)
		}
	}

	fmt.Printf("Primeiro vetor resultante : %d", paresComSoma)
	fmt.Printf("Segundo vetor resultante : %d", imparesComSoma)
}
