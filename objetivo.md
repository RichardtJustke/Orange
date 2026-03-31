# 🧠 1. Visão do Projeto

Antes de tudo, define isso aqui:

**Objetivo:**

> Criar uma aplicação CLI (terminal) para organizar notas, tarefas e informações, inspirada no Notion, com foco em produtividade e simplicidade.

**Usuário principal:**

* Você mesmo (dev / power user)
* Pessoas que gostam de terminal

---

# 🎯 2. Funcionalidades (Requisitos Funcionais)

Aqui é o que o sistema FAZ.

## 📄 Notas

* Criar nota
* Editar nota
* Deletar nota
* Listar notas
* Buscar nota (por título ou conteúdo)
* Tags nas notas

## 📋 Organização

* Criar “páginas” ou “workspaces”
* Hierarquia (tipo Notion mesmo):

  * Página → Subpágina → Nota
* Navegação entre páginas

## ✅ Tarefas

* Criar tarefas
* Marcar como concluída
* Listar tarefas pendentes/concluídas
* Prioridade (baixa, média, alta)

## 🔎 Busca

* Buscar global (notas + tarefas)
* Filtro por:

  * Tag
  * Data
  * Status

## 💾 Persistência

* Salvar dados localmente:

  * JSON (simples)
  * SQLite (mais avançado)

---

# ⚙️ 3. Requisitos Não Funcionais

Aqui é como o sistema se comporta:

* ⚡ Rápido (é CLI, tem que ser instantâneo)
* 🧩 Modular (facilitar adicionar features depois)
* 💻 Cross-platform (Linux, Windows, Mac)
* 📦 Offline first (sem depender de internet)
* 🧠 Fácil de usar via comandos

---

# 🖥️ 4. Interface (CLI UX)

Aqui é onde você se destaca 👀

Você pode seguir 2 estilos:

## 🔹 Estilo comando (tipo git)

```
app create note "Minha nota"
app list notes
app delete note 2
```

## 🔹 Estilo interativo (mais bonito)

* Menu navegável
* Setinhas ↑ ↓
* Enter pra selecionar

Exemplo:

```
> Notas
> Tarefas
> Buscar
> Sair
```

👉 Dica braba: usa libs tipo:

* `bubbletea` (Go) → UI linda no terminal
* `cobra` → CLI estruturada

---

# 🧱 5. Modelagem de Dados

Começa simples:

## Nota

```go
type Note struct {
    ID        int
    Title     string
    Content   string
    Tags      []string
    CreatedAt time.Time
}
```

## Tarefa

```go
type Task struct {
    ID        int
    Title     string
    Done      bool
    Priority  string
}
```

## Página

```go
type Page struct {
    ID       int
    Title    string
    ParentID *int
}
```

---

# 🏗️ 6. Arquitetura (importante pra não virar bagunça)

Organiza assim:

```
/cmd        -> entrada CLI
/internal
  /models   -> structs
  /services -> lógica (criar nota, etc)
  /storage  -> salvar dados
  /ui       -> interface terminal
```

---

# 🚀 7. Roadmap (pra não se perder)

Aqui é onde a maioria falha… então presta atenção:

## 🔹 Fase 1 (MVP)

* Criar nota
* Listar nota
* Salvar em JSON

## 🔹 Fase 2

* Editar / deletar
* Tags
* Busca

## 🔹 Fase 3

* Tarefas
* Filtros

## 🔹 Fase 4

* Interface interativa (menu bonito)

## 🔹 Fase 5 (modo monstro)

* Hierarquia tipo Notion
* SQLite
* Plugins ou extensões

---

# 🧠 8. Desafios que vão te fazer crescer

Esse projeto vai te forçar a aprender:

* Estrutura de CLI (cobra)
* Manipulação de arquivos
* Estrutura de dados real
* Arquitetura limpa
* UX no terminal (isso aqui é raro 👀)

---

# 💡 Ideias INSANAS (diferencial)

Se quiser sair do comum:

* 🔥 Markdown support
* 🧠 Autocomplete de comandos
* 🌙 Tema dark/light no terminal
* 🔗 Linkar notas (tipo Notion mesmo)
* 📊 Dashboard no terminal
* ⏱️ Timer estilo Pomodoro integrado

---
