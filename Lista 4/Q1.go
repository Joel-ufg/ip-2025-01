package main

import "fmt"

func main() {

	const tamanho = 10

	valores := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&valores[i])
	}

	encontrar := false
	for i := 0; i < tamanho; i++ {
		if valores[i] > 50 {
			fmt.Printf("Número %d na posição %d é maior que 50\n", valores[i], i)
			encontrar = true
		}
	}

	if !encontrar {
		fmt.Println("Nenhum número maior que 50 foi encontrado")
	}
}
