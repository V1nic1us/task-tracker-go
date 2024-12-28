# Task Trace Go

Leia este README em outros idiomas:
- [English](README.md)
- [Español](README.es.md)
- [Français](README.fr.md)

Task Tracker é um projeto para rastrear e gerenciar suas tarefas. Neste projeto, você construirá uma interface de linha de comando simples (CLI) para acompanhar o que você precisa fazer, o que já fez e no que está trabalhando no momento. Este projeto ajudará você a praticar suas habilidades de programação, incluindo trabalho com o sistema de arquivos, manipulação de entradas do usuário e construção de uma aplicação CLI simples.

## Requisitos

A aplicação deve ser executada na linha de comando, aceitar ações e entradas do usuário como argumentos e armazenar as tarefas em um arquivo JSON. O usuário deve ser capaz de:

- Adicionar, atualizar e excluir tarefas
- Marcar uma tarefa como em progresso ou concluída
- Listar todas as tarefas
- Listar todas as tarefas concluídas
- Listar todas as tarefas não concluídas
- Listar todas as tarefas em progresso

Aqui estão algumas restrições para orientar a implementação:

- Você pode usar qualquer linguagem de programação para construir este projeto.
- Use argumentos posicionais na linha de comando para aceitar entradas do usuário.
- Use um arquivo JSON para armazenar as tarefas no diretório atual.
- O arquivo JSON deve ser criado se não existir.
- Use o módulo de sistema de arquivos nativo da sua linguagem de programação para interagir com o arquivo JSON.
- Não use bibliotecas ou frameworks externos para construir este projeto.
- Certifique-se de lidar com erros e casos extremos de forma adequada.

## Exemplo

A lista de comandos e seus usos está descrita abaixo:

```bash
# Adicionando uma nova tarefa
task-cli add "Comprar mantimentos"
# Saída: Tarefa adicionada com sucesso (ID: 1)

# Atualizando e excluindo tarefas
task-cli update 1 "Comprar mantimentos e preparar o jantar"
task-cli delete 1

# Marcando uma tarefa como em progresso ou concluída
task-cli mark-in-progress 1
task-cli mark-done 1

# Listando todas as tarefas
task-cli list

# Listando tarefas por status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Propriedades de uma Tarefa

Cada tarefa deve ter as seguintes propriedades:

- **id**: Um identificador único para a tarefa.
- **description**: Uma breve descrição da tarefa.
- **status**: O status da tarefa (`todo`, `in-progress`, `done`).
- **createdAt**: A data e hora em que a tarefa foi criada.
- **updatedAt**: A data e hora da última atualização da tarefa.

Certifique-se de adicionar essas propriedades ao arquivo JSON ao adicionar uma nova tarefa e atualizá-las ao modificar uma tarefa.

## Como Começar

Aqui estão algumas etapas para ajudar você a iniciar o projeto **Task Tracker CLI**:

### Configuração do Ambiente de Desenvolvimento

1. Escolha uma linguagem de programação com a qual você esteja confortável (por exemplo, Python, JavaScript, etc.).
2. Certifique-se de ter um editor de código ou IDE instalado (por exemplo, VSCode, PyCharm).

### Inicialização do Projeto

1. Crie um novo diretório para o seu projeto **Task Tracker CLI**.
2. Inicialize um sistema de controle de versão (por exemplo, Git) para gerenciar o projeto.

### Implementação de Funcionalidades

1. Comece criando uma estrutura básica de CLI para manipular entradas do usuário.
2. Implemente cada funcionalidade uma por vez, garantindo que teste adequadamente antes de avançar para a próxima. Por exemplo:
   - Primeiro, implemente a funcionalidade de **adicionar tarefas**.
   - Em seguida, implemente **listar tarefas**, depois **atualizar tarefas**, e por fim **marcar tarefas como em progresso ou concluídas**.

### Testes e Depuração

1. Teste cada funcionalidade individualmente para garantir que funcionem como esperado.
2. Verifique o arquivo JSON para confirmar que as tarefas estão sendo armazenadas corretamente.
3. Depure quaisquer problemas que surgirem durante o desenvolvimento.

### Finalização do Projeto

1. Certifique-se de que todas as funcionalidades foram implementadas e testadas.
2. Limpe o código e adicione comentários onde necessário.
3. Escreva um arquivo **README** explicando como usar o seu **Task Tracker CLI**.

---

## Resultado Final

Ao final deste projeto, você terá desenvolvido uma ferramenta prática que pode ajudar você ou outras pessoas a gerenciar tarefas de maneira eficiente. Este projeto estabelece uma base sólida para projetos de programação mais avançados e aplicações do mundo real.

**Feliz codificação!**