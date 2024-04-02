package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 10
const delay = 5

func main() {

	for {

		exibeIntroducao()
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa!")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando!")
			os.Exit(-1)
		}

		fmt.Println(" ")
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibeIntroducao() {
	nome := "Cesar"
	versao := 1.1
	fmt.Println("Olá", nome, "!")
	fmt.Println("Este programa esta na versão", versao)
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"http://www.google.com.br", "http://www.bing.com.br"}
	sites := leSitesArquivo()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("--------------------------")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas! Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesArquivo() []string {
	var sites []string

	// arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")
	reader := bufio.NewReader(arquivo)

	if err != nil {
		fmt.Println("Erro", err)
	}

	for {
		linha, err := reader.ReadString('\n')

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro", err)
	}

	fmt.Println(string(arquivo))
}
