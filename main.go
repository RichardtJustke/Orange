package main

import "fmt"

// leitura de imput
// sistema de arquivos
type notes struct {
	ID      int
	Title   string
	Content string
}

var notes []Note

func main() {
	for {
		fmt.Println("\n1 - Criar nota")
		fmt.Println("2 - Listar notas")
		fmt.Println("3 - Sair")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createNote()
		case 2:
			listNotes()
		case 3:
			saveNotes()
			return
		}

	}
}
