# 📦 Guia Completo do `bufio` em Go

O pacote `bufio` fornece **operações de entrada e saída com buffer**, tornando leitura e escrita mais eficientes e flexíveis.

```go
import "bufio"
```

---

# 🧠 Conceito Principal

`bufio` funciona como uma camada entre você e o sistema:

👉 Em vez de ler/escrever direto (lento), ele usa um **buffer na memória** (rápido)

---

# 🔑 Principais Tipos

## 📥 1. `bufio.Reader`

Usado para leitura mais controlada (CLI, arquivos, streams)

### 🔹 Criar um Reader

```go
reader := bufio.NewReader(os.Stdin)
```

---

## 📚 Métodos do Reader

### 🟢 `ReadString(delim byte)`

Lê até encontrar um delimitador

```go
texto, err := reader.ReadString('\n')
```

---

### 🟢 `ReadBytes(delim byte)`

Igual ao `ReadString`, mas retorna `[]byte`

```go
data, err := reader.ReadBytes('\n')
```

---

### 🟢 `Read(p []byte)`

Lê bytes manualmente

```go
buffer := make([]byte, 10)
reader.Read(buffer)
```

---

### 🟢 `ReadByte()`

Lê um único byte

```go
b, err := reader.ReadByte()
```

---

### 🟢 `UnreadByte()`

“Deslê” o último byte

```go
reader.UnreadByte()
```

---

### 🟢 `Peek(n int)`

Espia sem consumir os dados

```go
data, _ := reader.Peek(5)
```

---

### 🟢 `Discard(n int)`

Ignora bytes

```go
reader.Discard(10)
```

---

# 📄 2. `bufio.Scanner`

Mais simples, ideal para leitura linha por linha

### 🔹 Criar um Scanner

```go
scanner := bufio.NewScanner(os.Stdin)
```

ou

```go
file, _ := os.Open("file.txt")
scanner := bufio.NewScanner(file)
```

---

## 📚 Métodos do Scanner

### 🟢 `Scan()`

Avança para o próximo token

```go
for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

---

### 🟢 `Text()`

Retorna o texto atual

```go
linha := scanner.Text()
```

---

### 🟢 `Bytes()`

Retorna como `[]byte`

```go
data := scanner.Bytes()
```

---

### 🟢 `Err()`

Verifica erros

```go
if err := scanner.Err(); err != nil {
	fmt.Println(err)
}
```

---

## 🔧 Split Functions (como ele divide os dados)

### 🟢 Padrão: linhas

```go
scanner.Split(bufio.ScanLines)
```

---

### 🟢 Outras opções:

```go
bufio.ScanWords      // palavras
bufio.ScanBytes      // byte por byte
bufio.ScanRunes      // caracteres unicode
```

---

### 🟢 Criar split customizado

```go
scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// lógica custom
})
```

---

# 📤 3. `bufio.Writer`

Usado para escrita com buffer

---

## 🔹 Criar um Writer

```go
writer := bufio.NewWriter(os.Stdout)
```

---

## 📚 Métodos do Writer

### 🟢 `Write(p []byte)`

```go
writer.Write([]byte("Hello"))
```

---

### 🟢 `WriteString(s string)`

```go
writer.WriteString("Hello\n")
```

---

### 🟢 `WriteByte(c byte)`

```go
writer.WriteByte('A')
```

---

### 🟢 `Flush()`

⚠️ **OBRIGATÓRIO**

```go
writer.Flush()
```

---

### 🟢 `Available()`

Mostra espaço disponível no buffer

```go
writer.Available()
```

---

### 🟢 `Buffered()`

Quantidade de dados no buffer

```go
writer.Buffered()
```

---

# ⚔️ Reader vs Scanner vs Writer

| Tipo    | Melhor uso                        |
| ------- | --------------------------------- |
| Reader  | Controle total da leitura         |
| Scanner | Leitura simples (linhas/palavras) |
| Writer  | Escrita eficiente                 |

---

# 🧪 Exemplo Completo (CLI simples)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite algo: ")
	texto, _ := reader.ReadString('\n')

	texto = strings.TrimSpace(texto)

	fmt.Println("Você digitou:", texto)
}
```

---

# 💡 Boas Práticas

✅ Use `Reader` para input de usuário
✅ Use `Scanner` para arquivos simples
✅ Sempre trate erros (`err`)
✅ Sempre use `Flush()` no Writer
✅ Use `TrimSpace` pra limpar entrada

---

# 🚨 Pegadinhas clássicas

* `Scanner` tem limite de buffer (~64KB por linha)
* `fmt.Scan` quebra com espaços → prefira `bufio`
* esquecer `Flush()` = nada aparece
* `\n` vem junto no input

---

# 🧠 Resumo Final (modo speedrun)

* `Reader` → leitura poderosa
* `Scanner` → leitura simples
* `Writer` → escrita rápida
* `bufio` = performance + controle

---
# 🧰 Kit básico do `bufio.Reader`

## 🔹 Criar o reader (sempre igual)

```go
reader := bufio.NewReader(os.Stdin)
```

---

## 🔹 Ler texto (padrão ouro)

```go
input, _ := reader.ReadString('\n')
```

👉 Isso lê **até o ENTER**

---

## 🔹 Limpar o input (ESSENCIAL)

```go
input = strings.TrimSpace(input)
```

👉 remove:

* `\n`
* espaços extras

Se você não fizer isso… seu programa começa a “agir estranho” 👀

---

# 🧠 Padrão que você vai usar 90% do tempo

```go
fmt.Print("Escolha uma opção: ")

input, _ := reader.ReadString('\n')
input = strings.TrimSpace(input)
```

Pronto. Isso aqui já resolve metade da vida.

---

# 🔢 Convertendo para número (menu CLI)

```go
import "strconv"

num, err := strconv.Atoi(input)

if err != nil {
	fmt.Println("Digite um número válido")
	return
}
```

---

# 🧪 Exemplo real (menu simples)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1 - Adicionar nota")
	fmt.Println("2 - Sair")

	fmt.Print("Escolha: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	opcao, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Opção inválida")
		return
	}

	if opcao == 1 {
		fmt.Print("Digite sua nota: ")

		nota, _ := reader.ReadString('\n')
		nota = strings.TrimSpace(nota)

		fmt.Println("Nota salva:", nota)
	} else {
		fmt.Println("Saindo...")
	}
}
```

---

# ⚠️ Erros clássicos (evita esses e você já tá acima da média)

❌ esquecer `TrimSpace`
❌ usar `fmt.Scan` (vai quebrar com espaço)
❌ ignorar erro de conversão
❌ misturar tipos sem converter

---

# 💡 Dica ninja (isso aqui muda o jogo)

Sempre crie uma função pra input:

```go
func lerInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
```

Uso:

```go
input := lerInput(reader)
```

👉 deixa seu código MUITO mais limpo

---

# 🧠 Resumo rápido

* `ReadString('\n')` → lê input
* `TrimSpace()` → limpa
* `Atoi()` → converte número
* repetir isso = CLI funcionando

