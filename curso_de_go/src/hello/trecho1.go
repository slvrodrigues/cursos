package main

import "fmt"

func main() {
	nome := "SlvRodrigues"
	idade := 24
	versao := 1.1
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
