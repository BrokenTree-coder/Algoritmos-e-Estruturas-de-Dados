package linkedlist

import (
	"testing"
)

var lists []ILinkedList // Usando a struct diretamente para evitar redeclarar a interface

func createLists() {
	linkedList := &LinkedList{}
	linkedList.Init()
	lists = []ILinkedList{linkedList}
}

func deleteLists() {
	lists = nil
}

func setupTest() func() {
	createLists()
	return func() {
		deleteLists()
	}
}

func TestAddAndGet(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// Adicionando elementos sequenciais
		for i := 0; i < 5; i++ {
			list.Add(i * 10)
		}

		if list.Size() != 5 {
			t.Errorf("Tamanho incorreto: obteve %d, esperava 5", list.Size())
		}

		// Verificando se os valores foram alocados na ordem certa
		for i := 0; i < 5; i++ {
			val, err := list.Get(i)
			if err != nil {
				t.Errorf("Erro inesperado no Get(%d): %v", i, err)
			}
			if val != i*10 {
				t.Errorf("Valor incorreto no índice %d: obteve %d, esperava %d", i, val, i*10)
			}
		}
	}
}

func TestGetOutOfBounds(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10) // index 0

		_, err := list.Get(1)
		if err == nil {
			t.Errorf("Esperava erro ao buscar índice inexistente")
		}

		_, err = list.Get(-1)
		if err == nil {
			t.Errorf("Esperava erro ao buscar índice negativo")
		}
	}
}

func TestAddOnIndex(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10)
		list.Add(30)
		// Estado atual: [10, 30]

		// Teste 1: Inserindo no meio (caminho feliz)
		err := list.AddOnIndex(20, 1)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		// Estado esperado: [10, 20, 30]

		// Teste 2: Inserindo no início (Pior cenário para ponteiros, altera o head)
		err = list.AddOnIndex(0, 0)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		// Estado esperado: [0, 10, 20, 30]

		// Teste 3: Inserindo no fim exato
		err = list.AddOnIndex(40, list.Size())
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		// Estado esperado: [0, 10, 20, 30, 40]

		expected := []int{0, 10, 20, 30, 40}
		if list.Size() != len(expected) {
			t.Errorf("Tamanho incorreto após AddOnIndex. Obteve %d, esperava %d", list.Size(), len(expected))
		}

		for i, exp := range expected {
			val, _ := list.Get(i)
			if val != exp {
				t.Errorf("No índice %d: obteve %d, esperava %d", i, val, exp)
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
			t.Errorf("Erro inesperado no Set: %v", err)
		}

		val, _ := list.Get(1)
		if val != 99 {
			t.Errorf("Falha no Set: obteve %d, esperava 99", val)
		}

		err = list.Set(100, 5) // Fora dos limites
		if err == nil {
			t.Errorf("Esperava erro ao usar Set fora dos limites")
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

		// Teste 1: Removendo do meio
		err := list.Remove(2)
		if err != nil {
			t.Errorf("Erro inesperado no Remove: %v", err)
		}
		// Estado esperado: [0, 10, 30, 40]

		val, _ := list.Get(2)
		if val != 30 {
			t.Errorf("Falha de ponteiros após Remove(meio). Índice 2 obteve %d, esperava 30", val)
		}

		// Teste 2: Removendo do início (Altera o head diretamente)
		list.Remove(0)
		// Estado esperado: [10, 30, 40]

		val, _ = list.Get(0)
		if val != 10 {
			t.Errorf("Falha ao remover o head. Índice 0 obteve %d, esperava 10", val)
		}

		// Teste 3: Removendo do fim
		list.Remove(list.Size() - 1)
		// Estado esperado: [10, 30]

		if list.Size() != 2 {
			t.Errorf("Tamanho incorreto após remoções. Obteve %d, esperava 2", list.Size())
		}
	}
}

func TestRemoveOutOfBounds(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(10)
		err := list.Remove(1) // Índice inválido para lista de tamanho 1
		if err == nil {
			t.Errorf("Esperava erro ao remover índice fora dos limites")
		}

		err = list.Remove(-1)
		if err == nil {
			t.Errorf("Esperava erro ao remover índice negativo")
		}
	}
}
