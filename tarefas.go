package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Variáveis globais para armazenar as tarefas
var (
	tarefas       [][]string                // Slice de slices: [título, descrição, status]
	statusTarefas = make(map[string]string) // Map para controle de status
	scanner       = bufio.NewScanner(os.Stdin)
)

// Função principal
func main() {
	fmt.Println("=== SISTEMA DE GERENCIAMENTO DE TAREFAS ===")
	fmt.Println()

	// Loop infinito para o menu
	for {
		exibirMenu()
		opcao := lerOpcao()

		switch opcao {
		case 1:
			adicionarTarefa()
		case 2:
			removerTarefa()
		case 3:
			marcarConcluida()
		case 4:
			listarTarefas()
		case 5:
			fmt.Println("Saindo do programa.")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}

		fmt.Println() // Linha em branco para melhor visualização
	}
}

// Exibe o menu principal
func exibirMenu() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1. Adicionar tarefa")
	fmt.Println("2. Remover tarefa")
	fmt.Println("3. Marcar tarefa como concluída")
	fmt.Println("4. Listar tarefas")
	fmt.Println("5. Sair")
	fmt.Print("Opção: ")
}

// Lê a opção do menu do usuário
func lerOpcao() int {
	scanner.Scan()
	opcaoStr := strings.TrimSpace(scanner.Text())

	opcao, err := strconv.Atoi(opcaoStr)
	if err != nil {
		return 0 // Retorna 0 se não conseguir converter
	}

	return opcao
}

// Lê uma linha de texto do usuário
func lerLinha() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Adiciona uma nova tarefa
func adicionarTarefa() {
	fmt.Print("Digite o título da tarefa: ")
	titulo := lerLinha()

	// Verifica se já existe uma tarefa com este título
	if _, existe := statusTarefas[titulo]; existe {
		fmt.Println("Já existe uma tarefa com este título!")
		return
	}

	fmt.Print("Digite a descrição da tarefa: ")
	descricao := lerLinha()

	// Cria nova tarefa como slice [título, descrição, status]
	novaTarefa := []string{titulo, descricao, "pendente"}

	// Adiciona a tarefa ao slice principal
	tarefas = append(tarefas, novaTarefa)

	// Adiciona o status ao map
	statusTarefas[titulo] = "pendente"

	fmt.Println("Tarefa adicionada com sucesso!")
}

// Remove uma tarefa existente
func removerTarefa() {
	if len(tarefas) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return
	}

	fmt.Print("Digite o título da tarefa a ser removida: ")
	titulo := lerLinha()

	// Encontra o índice da tarefa
	indice := encontrarTarefa(titulo)

	if indice == -1 {
		fmt.Println("Tarefa não encontrada.")
		return
	}

	// Remove a tarefa do slice (técnica Go para remoção)
	tarefas = append(tarefas[:indice], tarefas[indice+1:]...)

	// Remove do map de status
	delete(statusTarefas, titulo)

	fmt.Println("Tarefa removida com sucesso!")
}

// Marca uma tarefa como concluída
func marcarConcluida() {
	if len(tarefas) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return
	}

	fmt.Print("Digite o título da tarefa a ser marcada como concluída: ")
	titulo := lerLinha()

	// Encontra o índice da tarefa
	indice := encontrarTarefa(titulo)

	if indice == -1 {
		fmt.Println("Tarefa não encontrada.")
		return
	}

	// Atualiza o status na tarefa (posição 2 do slice)
	tarefas[indice][2] = "concluída"

	// Atualiza o status no map
	statusTarefas[titulo] = "concluída"

	fmt.Println("Tarefa marcada como concluída!")
}

// Lista todas as tarefas
func listarTarefas() {
	if len(tarefas) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return
	}

	fmt.Println("\n=== LISTA DE TAREFAS ===")

	for i, tarefa := range tarefas {
		titulo := tarefa[0]
		descricao := tarefa[1]
		status := tarefa[2]

		fmt.Printf("%d. Título: %s\n", i+1, titulo)
		fmt.Printf("   Descrição: %s\n", descricao)
		fmt.Printf("   Status: %s\n", status)
		fmt.Println("   " + strings.Repeat("-", 30))
	}
}

// Função auxiliar para encontrar uma tarefa pelo título
func encontrarTarefa(titulo string) int {
	for i, tarefa := range tarefas {
		if tarefa[0] == titulo { // tarefa[0] é o título
			return i
		}
	}
	return -1 // Não encontrada
}
