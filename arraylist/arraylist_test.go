package arraylist

import (
	"testing"
)

var initialCapacity int
var lists []IArrayList

// createLists inicializa as estruturas que serão testadas
func createLists(capacity int) {
	arrayList := &ArrayList{}
	(*arrayList).Init(capacity)

	// No futuro, será adicionada a LinkedList aqui:
	// linkedList := &LinkedList{}
	// lists = []IArrayList{arrayList, linkedList}

	lists = []IArrayList{arrayList}
}

func deleteLists() {
	lists = nil
}

func setupTest() func() {
	initialCapacity = 5
	createLists(initialCapacity)

	return func() {
		deleteLists()
	}
}

func TestAddAndGet(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// Adiciona o dobro da capacidade para forçar o resize()
		for i := 0; i < initialCapacity*2; i++ {
			list.Add(i * 10)
		}

		if list.Size() != initialCapacity*2 {
			t.Errorf("%T size = %d, expected %d", list, list.Size(), initialCapacity*2)
		}

		// Verifica se os valores estão corretos e na ordem certa
		for i := 0; i < initialCapacity*2; i++ {
			val, err := list.Get(i)
			if err != nil {
				t.Errorf("%T unexpected error: %v", list, err)
			}
			if val != i*10 {
				t.Errorf("%T got %d, expected %d", list, val, i*10)
			}
		}
	}
}

func TestGetOutOfBounds(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10) // index 0

		_, err := list.Get(1) // Tentando acessar fora dos limites
		if err == nil {
			t.Errorf("%T expected error on out of bounds get", list)
		}

		_, err = list.Get(-1) // Índice negativo
		if err == nil {
			t.Errorf("%T expected error on negative index get", list)
		}
	}
}

func TestAddOnIndex(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10)
		list.Add(30)
		// Estado atual: [10, 30]

		// Inserindo no meio
		err := list.AddOnIndex(20, 1)
		if err != nil {
			t.Errorf("%T unexpected error: %v", list, err)
		}
		// Estado esperado: [10, 20, 30]

		// Inserindo no início (Pior caso O(n))
		err = list.AddOnIndex(0, 0)
		if err != nil {
			t.Errorf("%T unexpected error: %v", list, err)
		}
		// Estado esperado: [0, 10, 20, 30]

		// Inserindo no fim (Melhor caso O(1))
		err = list.AddOnIndex(40, list.Size())
		if err != nil {
			t.Errorf("%T unexpected error: %v", list, err)
		}
		// Estado esperado: [0, 10, 20, 30, 40]

		expected := []int{0, 10, 20, 30, 40}
		if list.Size() != len(expected) {
			t.Errorf("%T wrong size after AddOnIndex", list)
		}

		for i, exp := range expected {
			val, _ := list.Get(i)
			if val != exp {
				t.Errorf("%T at index %d got %d, expected %d", list, i, val, exp)
			}
		}
	}
}

func TestSet(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10)
		list.Add(20)

		err := list.Set(99, 1)
		if err != nil {
			t.Errorf("%T unexpected error: %v", list, err)
		}

		val, _ := list.Get(1)
		if val != 99 {
			t.Errorf("%T set failed, got %d expected 99", list, val)
		}

		err = list.Set(100, 5) // Fora dos limites
		if err == nil {
			t.Errorf("%T expected error when setting out of bounds", list)
		}
	}
}

func TestRemove(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		for i := 0; i < 5; i++ {
			list.Add(i * 10)
		}
		// Estado atual: [0, 10, 20, 30, 40]

		// Removendo do meio
		err := list.Remove(2)
		if err != nil {
			t.Errorf("%T unexpected error: %v", list, err)
		}
		// Estado esperado: [0, 10, 30, 40]

		val, _ := list.Get(2)
		if val != 30 {
			t.Errorf("%T shift failed, got %d expected 30", list, val)
		}

		// Removendo do início (Pior caso O(n))
		list.Remove(0)
		// Estado esperado: [10, 30, 40]

		val, _ = list.Get(0)
		if val != 10 {
			t.Errorf("%T shift failed, got %d expected 10", list, val)
		}

		// Removendo do fim (Melhor caso O(1))
		list.Remove(list.Size() - 1)
		// Estado esperado: [10, 30]

		if list.Size() != 2 {
			t.Errorf("%T wrong size after removals, got %d expected 2", list, list.Size())
		}
	}
}

func TestRemoveOutOfBounds(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10)
		err := list.Remove(1) // Índice inválido (só existe o 0)
		if err == nil {
			t.Errorf("%T expected error when removing out of bounds", list)
		}
	}
}
