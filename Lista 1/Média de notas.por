programa {
  
  inclua biblioteca Matematica --> mat
  
  funcao inicio() {
    cadeia nome
    real nota1, nota2, nota3, media

    escreva("Digite seu nome: ")
		leia(nome)
    
    escreva("/n")

    escreva("Primeira nota: ")
    leia(nota1)

    escreva("Segunda nota: ")
    leia(nota2)

escreva("Terceira nota: ")
leia(nota3)

/* Calcula a média final do usuário */
		media = (nota1 + nota2 + nota3) / 3

    limpa()

    se (media>=6 )
    {
		 	escreva("Parabéns ", nome, "!\nVocê foi aprovado com a média ", mat.arredondar(media, 2))
		}
		senao
		{
			escreva("Que pena ", nome, "!\nVocê foi reprovado com a média ", mat.arredondar(media, 2))
		}
  }
}
