package doublelinkedlist

import (
	"testing"
)

var lists []*DoublyLinkedList

func createLists() {
	dll := &DoublyLinkedList{}
	dll.Init()
	lists = []*DoublyLinkedList{dll}
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
		// Adiciona 5 elementos (Testa a atualização constante do tail)
		for i := 0; i < 5; i++ {
			list.Add(i * 10)
		}

		if list.Size() != 5 {
			t.Errorf("Tamanho incorreto: obteve %d, esperava 5", list.Size())
		}

		// Testa o Get (Que internamente vai testar a sua otimização de busca bidirecional)
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

func TestAddOnIndex(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// Teste 1: Inserindo no index 0 com a lista VAZIA (Testa se head e tail são criados juntos)
		err := list.AddOnIndex(30, 0)
		if err != nil {
			t.Errorf("Erro inesperado ao inserir em lista vazia: %v", err)
		}

		// Teste 2: Inserindo no index 0 com a lista populada (O head recua)
		list.AddOnIndex(10, 0)

		// Teste 3: Inserindo no meio (A cirurgia de 4 setas)
		list.AddOnIndex(20, 1)

		// Teste 4: Inserindo no fim exato (O tail avança)
		list.AddOnIndex(40, list.Size())

		// Estado esperado: [10, 20, 30, 40]
		expected := []int{10, 20, 30, 40}
		if list.Size() != len(expected) {
			t.Errorf("Tamanho incorreto após múltiplos AddOnIndex. Obteve %d, esperava %d", list.Size(), len(expected))
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
		list.Add(30) // Estado: [10, 20, 30]

		// Set no meio para garantir que não rompe os ponteiros
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
		// Preparando uma lista com 5 elementos: [0, 10, 20, 30, 40]
		for i := 0; i < 5; i++ {
			list.Add(i * 10)
		}

		// Teste 1: Removendo do meio (Testa o anterior.next = posterior, etc)
		err := list.Remove(2) // Remove o 20
		if err != nil {
			t.Errorf("Erro inesperado no Remove: %v", err)
		}

		// Teste 2: Removendo do início (Altera o head diretamente)
		list.Remove(0) // Remove o 0

		// Teste 3: Removendo do fim (Altera o tail diretamente)
		list.Remove(list.Size() - 1) // Remove o 40

		// Estado esperado restante: [10, 30]
		if list.Size() != 2 {
			t.Errorf("Tamanho incorreto após remoções. Obteve %d, esperava 2", list.Size())
		}

		val, _ := list.Get(0)
		if val != 10 {
			t.Errorf("Falha de ponteiros no head após remoções. Obteve %d, esperava 10", val)
		}

		val, _ = list.Get(1)
		if val != 30 {
			t.Errorf("Falha de ponteiros no tail após remoções. Obteve %d, esperava 30", val)
		}
	}
}

func TestRemoveOnlyElement(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		list.Add(99) // Lista de tamanho 1

		err := list.Remove(0)
		if err != nil {
			t.Errorf("Erro inesperado ao remover o único elemento: %v", err)
		}

		if list.Size() != 0 {
			t.Errorf("Tamanho deveria ser 0, obteve %d", list.Size())
		}

		// Tentar pegar algo deve dar erro agora
		_, err = list.Get(0)
		if err == nil {
			t.Errorf("A lista deveria estar vazia, mas um elemento foi encontrado")
		}
	}
}
