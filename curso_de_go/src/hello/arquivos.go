package main

import (
	"fmt" //pacote responsavel por fazer requisições web
	"net/http"
	"os"   //pacote responsavel por informar a saída do comando ao sistema, status 0 1
	"time" //pacote com função sleep temporizador
)

const monitoramentos = 3 //variável constante com o número de repetições
const delay = 5          //variável constante com o tempo em segundos para o intervalo entre monitoramentos

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
	fmt.Println("")

	return comando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://www.globo.com", "https://www.terra.com.br", "https://www.digipix.com.br/tetao", "https://www.hotmail.com"}

	//fmt.Println(sites)
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ { //esse for irá executar 5 testes consecutivos a variavel monitoramentos está recebendo o valor

		for i, site := range sites { //range retorna 2 coisas, índice e item  da posição
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second) //tempo pode ser definido desde milisegundos até minutos declarado na constante no inicio do codigo
		fmt.Println("")
	}

}

func testaSite(site string) { //função criada com retorno de status code
	resp, err := http.Get(site) //resposta e erro solicitados em retorno

	if err != nil { //se erro é diferente de nulo, s recebe o erro contendo o seu código
		s := err.Error()
		fmt.Printf("%q\n", s)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, _ := os.Open("sites.txt") //abre um arquivo e retorna
	fmt.Println(arquivo)
	return sites
}
