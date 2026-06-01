# 💻 Algoritmos e Estruturas de Dados (em Go)

Repositório dedicado à implementação de algoritmos e estruturas de dados fundamentais, desenvolvidos inteiramente em Go (Golang).

Este repositório está sendo construído de forma contínua durante a disciplina de Algoritmos e Estruturas de Dados do curso de Engenharia da Computação na Universidade Federal do Rio Grande do Norte (UFRN).

## 📂 Estruturas e Algoritmos Implementados

Até o momento, o repositório conta com as seguintes implementações, divididas por diretórios:

### 📝 Listas (Lists)

- ArrayList: Implementação de lista baseada em arrays dinâmicos.

- LinkedList: Lista encadeada simples (Single Linked List).

- DoubleLinkedList: Lista duplamente encadeada (Doubly Linked List).

### 📚 Pilhas (Stacks)

- StackArray: Pilha implementada utilizando arrays estáticos/dinâmicos sob o capô.

- StackLinkedList: Pilha implementada baseada em nós e ponteiros (Lista Encadeada).

### 🚶 Filas (Queues)

- QueueFifo: Fila padrão com comportamento First-In-First-Out.

- Deque: Fila Duplamente Terminada (Double-Ended Queue), permitindo inserção e remoção em ambas as extremidades.

### 🔍 Algoritmos de Busca (Search Algorithms)

- SearchAlgorithms: Implementações de buscas clássicas (como Busca Linear e Busca Binária).

## 🛠️ Como Executar

Como essas implementações foram feitas em Go, certifique-se de ter o Go instalado em sua máquina.

- Clone o repositório:

```bash
git clone [https://github.com/BrokenTree-coder/Algoritmos-e-Estruturas-de-Dados.git](https://github.com/BrokenTree-coder/Algoritmos-e-Estruturas-de-Dados.git)
```

- Acesse o diretório do projeto:

```bash
cd Algoritmos-e-Estruturas-de-Dados
```

- Navegue até a estrutura que deseja testar e execute os arquivos principais ou de teste:

```bash
# Exemplo rodando testes da lista encadeada (caso existam arquivos _test.go)

cd linkedlist

go test -v
```
