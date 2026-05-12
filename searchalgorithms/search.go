package searchalgorithms

// busca linear:
// eu pensei nessa implementação baseando-se nos meus conhecimentos anteriores de loop em outras linguagens de programação
// func LinearSearch(elementos []int, val int) int {
// 	for i := 0; i < len(elementos); i++ {
// 		if elementos[i] == val {
// 			return i
// 		}
// 	}
//
// 	return -1
// }

// essa é a implementação idiomática/endêmica da linguagem Go, usando o 'range'
// quando você usa 'for i, v := range elementos'
// o Go faz uma cópia do valor do elemento atual para a variável v em cada iteração, e o indice em i
// como padrão, usar o range pela segurança e clareza de código, ele é mais limpo e evita erros de índice
// só mudar para o loop clássico (ou usar 'for i := range elementos' sem copiar o valor
// se você estiver lidando com tipos de dados muito grandes onde a cópia se torne um gargalo

func LinearSearch(elementos []int, val int) int {
	for i, v := range elementos {
		if v == val {
			return i
		}
	}
	return -1
}

// busca binária:
// o array obrigatoriamente deve estar ordenado, pra essa busca fucionar
// a complexidade dessa busca é bem menor que a linear
// linear --> O(n)
// binária --> O(log(n))
// ex: se o array tiver 1024 elementos
// O(n) --> a linear pode ter que percorrer o array inteiro, ou seja, dar 1024 passos
// O(log(n)) --> a binária levará no máximo 10 passos [ log(1024) = 10 ]

func BinarySearch(elementos []int, val int) int {
	inicio := 0
	fim := len(elementos) - 1

	for inicio <= fim {
		// calcula o meio uma vez por iteração
		meio := inicio + (fim-inicio)/2

		if elementos[meio] == val {
			return meio // encontrou! retorna o índice.
		}

		if val > elementos[meio] {
			// se o valor é maior, descarta o meio e tudo à esquerda
			inicio = meio + 1
		} else {
			// se o valor é menor, descarta o meio e tudo à direita
			fim = meio - 1
		}
	}

	return -1
}

// elementos = [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// BinarySearch(elementos, 7)

// iteracao:
// [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
//                  i     m     f

// iteracao:
// val == elementos[meio]

// elementos = [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// BinarySearch(elementos, 0)

// iteracao:
// [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
//   i  m     f

// iteracao:
// [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
//   i
//   f
//   m

// iteracao:
// val == elementos[meio]
