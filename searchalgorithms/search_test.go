package searchalgorithms

import "testing"

func TestSearchAlgorithms(t *testing.T) {
	// O array para busca binária PRECISA estar ordenado
	elementos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	testes := []struct {
		nome     string
		alvo     int
		esperado int
	}{
		{"Encontrar no início", 0, 0},
		{"Encontrar no meio", 4, 4},
		{"Encontrar no fim", 9, 9},
		{"Elemento inexistente", 15, -1},
	}

	for _, tt := range testes {
		t.Run(tt.nome+"_Linear", func(t *testing.T) {
			res := LinearSearch(elementos, tt.alvo)
			if res != tt.esperado {
				t.Errorf("LinearSearch(%d) = %d; esperado %d", tt.alvo, res, tt.esperado)
			}
		})

		t.Run(tt.nome+"_Binary", func(t *testing.T) {
			res := BinarySearch(elementos, tt.alvo)
			if res != tt.esperado {
				t.Errorf("BinarySearch(%d) = %d; esperado %d", tt.alvo, res, tt.esperado)
			}
		})
	}
}
