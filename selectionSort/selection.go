package selectionSort

import (
	"math"
)

// SelectionSorter é a estrutura que implementa a interface ISorter.
type SelectionSorter struct{}

// NewSelectionSorter atua como um construtor, retornando a interface ISorter.
func NewSelectionSorter() ISelection {
	return &SelectionSorter{}
}

func (s *SelectionSorter) SelectionSort_OutOfPlace(v []int) []int {
	tamanho := len(v)
	ord := make([]int, tamanho)

	// criando uma cópia do array original pra evitar corromper o slice original de entrada
	desord := make([]int, tamanho)
	copy(desord, v)

	for i := 0; i < tamanho; i++ {
		menor := 0

		for j := 1; j < tamanho; j++ {
			if desord[j] < desord[menor] {
				menor = j
			}
		}

		ord[i] = desord[menor]

		// "substituiremos seu valor no Array Original pelo maior valor de int" (Valor Sentinela)
		desord[menor] = math.MaxInt
	}

	return ord
}

func (s *SelectionSorter) SelectionSort_InPlace(v []int) []int {
	// só precisa ir até o penultimo, por isso, o 'for' só vai até --> 'len(v) - 1'
	for i := 0; i < len(v)-1; i++ {
		menor := i
		for j := i + 1; j < len(v); j++ {
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[i], v[menor] = v[menor], v[i]
	}

	return v
}
