package main

import (
	"fmt" //pacote responsavel por fazer requisições web
	"net/http"
	"os" //pacote responsavel por informar a saída do comando ao sistema, status 0 1
	"strings"
)

func main() {

	exibeIntroducao()

	for { //A linguagem go não possuí while, dessa forma o programa entrara em loop usando for

		exibeMenu()

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
	var site string
	fmt.Println("digite um site acrescido de https://")
	fmt.Scanln(&site)
	if !strings.HasPrefix(site, "https") {
		site = fmt.Sprintf("https://%s", site)
	}
	fmt.Println("Monitorando...")
	//fmt.Scan("digite um site", &site)
	//site := "https://www.alura.com.br"
	resp, _ := http.Get(site) //função de retorno , status code http 200 OK diferente disso retorna msg de erro
	fmt.Println(resp)         //retorna resposta em tela com cabeçalho de acesso

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")

	} else {
		fmt.Println("Site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}
