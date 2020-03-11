package main

import (
	"fmt" //pacote responsavel por fazer requisições web
	"net/http"
	"os"      //pacote responsavel por informar a saída do comando ao sistema, status 0 1
	"reflect" //pacote responsável por descobrir o tipo typeof
)

func main() {

	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://www.globo.com"
	sites[1] = "https://www.terra.com.br"
	sites[2] = "https://www.caelum.com.br"
	sites[3] = "https://www.hotmail.com"
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)

	//exibeIntroducao()

	for { //A linguagem go não possuí while, dessa forma o programa entrara em loop usando for

		//exibeMenu()
		exibeNomes()

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
	var sites [4]string
	sites[0] = "https://www.globo.com"
	sites[1] = "https://www.terra.com.br"
	sites[2] = "https://www.caelum.com.br"
	sites[3] = "https://www.hotmail.com"

	fmt.Println(sites)

	site := "https://r7.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")

	} else {
		fmt.Println("Site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"} //não é necessário declarar tamanho em um slice
	//nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")                  // len conta as posições contidas dentro do slice
	fmt.Println("O meu slice tem capacidade para ", cap(nomes), "itens") // cap retorna a capacidade do slice

}
