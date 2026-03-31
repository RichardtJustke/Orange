package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite seu nome:")
	nome, _ := reader.ReadString('\n')

	fmt.Println("Olá,", nome)
}
