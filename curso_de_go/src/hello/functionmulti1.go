package main

import (
	"fmt" //pacote responsavel por fazer requisições web
	"os"  //pacote responsavel por informar a saída do comando ao sistema, status 0 1
)

func main() {

	//exibeIntroducao()
	//exibeMenu()

	nome, idade := devolveNomeEIdade()
	fmt.Println(nome, "E tenho", idade, "anos")

	comando := leComando() //comando será o retorno da função leComando

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0) //saída com sucesso
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1) //saída com tratamento de erro
	}
}

func devolveNomeEIdade() (string, int) { //Funcao que devolve multiplos retornos
	nome := "Emerson"
	idade := 24
	return nome, idade
}

func exibeIntroducao() { //funcao que não retorna nada, apenas exibe
	nome := "Slvrodrigues"
	idade := 34
	versao := 1.1
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//site := "https://www.alura.com.br"
	//resp, err:= http.Get(site)
}
