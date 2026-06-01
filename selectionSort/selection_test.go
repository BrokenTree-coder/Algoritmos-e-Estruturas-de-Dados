package selectionSort

import (
	"reflect"
	"testing"
)

func TestSelectionSortOutOfPlace(t *testing.T) {
	// Instancia o  ordenador usando a interface
	sorter := NewSelectionSorter()

	// Define os cenários de teste (Table-Driven Tests)
	cenarios := []struct {
		nome     string
		entrada  []int
		esperado []int
	}{
		{
			nome:     "Vetor desordenado normal",
			entrada:  []int{8, 2, 4, 3, 7},
			esperado: []int{2, 3, 4, 7, 8},
		},
		{
			nome:     "Vetor já ordenado",
			entrada:  []int{1, 2, 3, 4, 5},
			esperado: []int{1, 2, 3, 4, 5},
		},
		{
			nome:     "Vetor com elementos repetidos",
			entrada:  []int{5, 1, 5, 2, 1},
			esperado: []int{1, 1, 2, 5, 5},
		},
		{
			nome:     "Vetor com números negativos",
			entrada:  []int{3, -1, -4, 2, 0},
			esperado: []int{-4, -1, 0, 2, 3},
		},
		{
			nome:     "Vetor vazio",
			entrada:  []int{},
			esperado: []int{},
		},
	}

	// Executa os cenários
	for _, cenario := range cenarios {
		t.Run(cenario.nome, func(t *testing.T) {
			resultado := sorter.SelectionSort_OutOfPlace(cenario.entrada)

			// reflect.DeepEqual é usado para comparar slices em Go
			if !reflect.DeepEqual(resultado, cenario.esperado) {
				t.Errorf("Falha no cenário '%s'. \nEsperado: %v \nObtido: %v", cenario.nome, cenario.esperado, resultado)
			}
		})
	}
}

func TestSelectionSortInPlace(t *testing.T) {
	sorter := NewSelectionSorter()

	cenarios := []struct {
		nome     string
		entrada  []int
		esperado []int
	}{
		{
			nome:     "Vetor desordenado normal",
			entrada:  []int{8, 2, 4, 3, 7},
			esperado: []int{2, 3, 4, 7, 8},
		},
		{
			nome:     "Vetor já ordenado",
			entrada:  []int{1, 2, 3, 4, 5},
			esperado: []int{1, 2, 3, 4, 5},
		},
		{
			nome:     "Vetor com elementos repetidos",
			entrada:  []int{5, 1, 5, 2, 1},
			esperado: []int{1, 1, 2, 5, 5},
		},
		{
			nome:     "Vetor com números negativos",
			entrada:  []int{3, -1, -4, 2, 0},
			esperado: []int{-4, -1, 0, 2, 3},
		},
		{
			nome:     "Vetor vazio",
			entrada:  []int{},
			esperado: []int{},
		},
	}

	for _, cenario := range cenarios {
		t.Run(cenario.nome, func(t *testing.T) {
			resultado := sorter.SelectionSort_InPlace(cenario.entrada)

			if !reflect.DeepEqual(resultado, cenario.esperado) {
				t.Errorf("Falha no cenário '%s'. \nEsperado: %v \nObtido: %v", cenario.nome, cenario.esperado, resultado)
			}
		})
	}
}
