package main

import (
	"bufio"
	"fmt" //pacote responsavel por fazer requisições web
	"io"
	"net/http"
	"os" //pacote responsavel por informar a saída do comando ao sistema, status 0 1
	"strings"
	"time" //pacote com função sleep temporizador
)

const monitoramentos = 3 //variável constante com o número de repetições
const delay = 5          //variável constante com o tempo em segundos para o intervalo entre monitoramentos

func main() {
	exibeIntroducao()
	leSitesDoArquivo()
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

	// criando o slice a partir da função leSitesDoArquivo()
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
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
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}

	leitor := bufio.NewReader(arquivo) //pacote bufio percorre linha a linha do arquivo
	for {                              //for infinito para ler linha a linha o arquivo txt

		linha, err := leitor.ReadString('\n') //aspas duplas não pode ser usado pois retorna uma string
		linha = strings.TrimSpace(linha)      //remove o espaço em branco
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {

			break
		}

	}

	//fmt.Println(sites)
	arquivo.Close()
	return sites

}
