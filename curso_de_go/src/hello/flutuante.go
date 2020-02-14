package main

import "fmt"

func main() {
	var precoLeite float32 = 2.99
	var precoOvo float32 = 3.99
	var precoPao float32 = 0.99

	fmt.Println("O preço do leite é R$", precoLeite)
	fmt.Println("O preço do ovo é R$", precoOvo)
	fmt.Println("O preço do pão é R$", precoPao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando)
}
