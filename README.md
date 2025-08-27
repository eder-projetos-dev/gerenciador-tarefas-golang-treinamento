# Gerenciador de Tarefas - Golang Treinamento

Sistema de gerenciamento de tarefas desenvolvido em Go como exercício de aprendizado dos fundamentos da linguagem.

## 📋 Sobre o Projeto

Este é um sistema de linha de comando que permite gerenciar tarefas pessoais através de um menu interativo. O projeto faz parte de uma série de exercícios práticos para aprendizado da linguagem Go.

## ✨ Funcionalidades

- ➕ **Adicionar tarefa**: Criar nova tarefa com título e descrição
- ❌ **Remover tarefa**: Excluir tarefa existente pelo título
- ✅ **Marcar como concluída**: Alterar status de tarefa para "concluída"
- 📝 **Listar tarefas**: Visualizar todas as tarefas com seus status
- 🚪 **Sair**: Encerrar o programa

## 🛠️ Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Interface**: Terminal/CLI
- **Estruturas de dados**: Slices e Maps

## 📚 Conceitos de Go Praticados

- Declaração de variáveis e tipos básicos
- Slices e manipulação de arrays dinâmicos
- Maps (mapas/dicionários)
- Estruturas de controle (loops, switch/case)
- Funções e organização de código
- Entrada e saída de dados no terminal
- Tratamento de strings

## 🚀 Como Executar

### Pré-requisitos
- Go instalado (versão 1.19 ou superior)

### Executando o programa
```bash
# Clone o repositório
git clone https://github.com/seu-usuario/gerenciador-tarefas-golang-treinamento.git

# Navegue até o diretório
cd gerenciador-tarefas-golang-treinamento

# Execute o programa
go run tarefas.go
```

### Alternativa - Compilar e executar
```bash
# Compile o programa
go build tarefas.go

# Execute o binário (Linux/Mac)
./tarefas

# Execute o binário (Windows)
tarefas.exe
```

## 🎮 Como Usar

1. Execute o programa
2. Escolha uma opção do menu (1-5)
3. Siga as instruções na tela
4. Para sair, escolha a opção 5

### Exemplo de uso:
```
=== SISTEMA DE GERENCIAMENTO DE TAREFAS ===

Escolha uma opção:
1. Adicionar tarefa
2. Remover tarefa
3. Marcar tarefa como concluída
4. Listar tarefas
5. Sair
Opção: 1

Digite o título da tarefa: Comprar leite
Digite a descrição da tarefa: Comprar 2 litros de leite no supermercado
Tarefa adicionada com sucesso!
```

## 📁 Estrutura do Projeto

```
gerenciador-tarefas-golang-treinamento/
├── README.md                    # Documentação do projeto
├── tarefas.go                   # Código fonte principal
├── requisitos/                  # Documentação técnica
│   └── Exercicio-Integrador-Lista-de-Tarefas.pdf
└── exemplos/                    # Exemplos de execução
    └── execucao-demo.txt
```

## 🎯 Objetivos de Aprendizado

Este exercício foi desenvolvido para praticar:

- [x] Estruturação de programas Go
- [x] Uso de slices e manipulação de dados
- [x] Implementação de maps para controle de estado
- [x] Criação de funções organizadas e reutilizáveis
- [x] Interação com usuário via terminal
- [x] Validação de entrada e tratamento de erros
- [x] Loops e estruturas condicionais

## 📋 Requisitos Funcionais Implementados

- ✅ Menu principal interativo
- ✅ Adicionar tarefa com título e descrição
- ✅ Remover tarefa por título
- ✅ Marcar tarefa como concluída
- ✅ Listar todas as tarefas
- ✅ Validação de entradas do usuário
- ✅ Mensagens informativas para todas as operações

## 🔧 Detalhes Técnicos

### Estruturas de Dados
- **Slice de slices**: `[][]string` - Armazena [título, descrição, status]
- **Map**: `map[string]string` - Controle rápido de status por título

### Arquitetura
- Variáveis globais para estado das tarefas
- Funções separadas para cada operação do menu
- Função auxiliar para busca de tarefas
- Scanner para leitura segura de entrada do usuário

## 🤝 Contribuições

Este é um projeto educacional, mas sugestões e melhorias são bem-vindas!

## 📄 Licença

Este projeto é parte de um treinamento educacional e está disponível para fins de aprendizado.

---

**Desenvolvido durante o treinamento de Golang** 🐹