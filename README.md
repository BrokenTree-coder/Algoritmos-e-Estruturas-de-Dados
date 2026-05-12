\# 💻 Algoritmos e Estruturas de Dados (em Go)



Repositório dedicado à implementação de algoritmos e estruturas de dados fundamentais, desenvolvidos inteiramente em \*\*Go (Golang)\*\*. 



Este projeto está sendo construído de forma contínua durante a disciplina de Algoritmos e Estruturas de Dados do curso de Engenharia da Computação na Universidade Federal do Rio Grande do Norte (UFRN).



\---



\## 📌 Status do Projeto

🚧 \*\*Em desenvolvimento:\*\* Novos módulos, otimizações e estruturas serão adicionados progressivamente conforme o avanço dos tópicos da disciplina ao longo do semestre.



\---



\## 📂 Estruturas e Algoritmos Implementados



Até o momento, o repositório conta com as seguintes implementações, divididas por diretórios:



\### 📝 Listas (Lists)

\- \*\*\[`arraylist/`](./arraylist):\*\* Implementação de lista baseada em arrays dinâmicos.

\- \*\*\[`linkedlist/`](./linkedlist):\*\* Lista encadeada simples (Single Linked List).

\- \*\*\[`doublelinkedlist/`](./doublelinkedlist):\*\* Lista duplamente encadeada (Doubly Linked List).



\### 📚 Pilhas (Stacks)

\- \*\*\[`stackarray/`](./stackarray):\*\* Pilha implementada utilizando arrays estáticos/dinâmicos sob o capô.

\- \*\*\[`stacklinkedlist/`](./stacklinkedlist):\*\* Pilha implementada baseada em nós e ponteiros (Lista Encadeada).



\### 🚶 Filas (Queues)

\- \*\*\[`queuefifo/`](./queuefifo):\*\* Fila padrão com comportamento First-In-First-Out.

\- \*\*\[`deque/`](./deque):\*\* Fila Duplamente Terminada (Double-Ended Queue), permitindo inserção e remoção em ambas as extremidades.



\### 🔍 Algoritmos de Busca (Search Algorithms)

\- \*\*\[`searchalgorithms/`](./searchalgorithms):\*\* Implementações de buscas clássicas (como Busca Linear e Busca Binária).



\---



\## 🚀 Próximos Passos (Roadmap)

Conforme a disciplina avançar, planejo implementar:

\- \[ ] Árvores (Binárias, AVL, Red-Black)

\- \[ ] Algoritmos de Ordenação (MergeSort, QuickSort, etc.)

\- \[ ] Tabelas Hash (Hash Tables)

\- \[ ] Grafos (Busca em Largura, Busca em Profundidade)



\---



\## 🛠️ Como Executar



Como este projeto é um módulo Go, certifique-se de ter o Go instalado em sua máquina.



1\. Clone o repositório:



```bash

git clone \[https://github.com/BrokenTree-coder/Algoritmos-e-Estruturas-de-Dados.git](https://github.com/BrokenTree-coder/Algoritmos-e-Estruturas-de-Dados.git)





2\. Acesse o diretório do projeto:



```bash

cd Algoritmos-e-Estruturas-de-Dados





3\. Navegue até a estrutura que deseja testar e execute os arquivos principais ou de teste:



```bash

\# Exemplo rodando testes da lista encadeada (caso existam arquivos \_test.go)

cd linkedlist

go test -v





