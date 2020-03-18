package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	InputDoUtilizador := bufio.NewReader(os.Stdin)

	fmt.Print("insere o teu primeiro nome:")
	primeiroNome, _ := InputDoUtilizador.ReadString('\n')

	fmt.Print("insere o teu último nome:")
	ultimoNome, _ := InputDoUtilizador.ReadString('\n')

	fmt.Print("insere a tua idade:")
	idade, _ := InputDoUtilizador.ReadString('\n')

	fmt.Printf("Primeiro nome: %sUltimo nome: %sIdade: %s", primeiroNome, ultimoNome, idade)

}
