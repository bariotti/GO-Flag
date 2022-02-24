package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	path := flag.String("path", "myapp.log", "Determina o arquivo de LOG que será lido e vasculhado")
	logType := flag.String("logType", "ERRO", "Determina o tipo de LOG que será vasculhado no arquivo. Opções: ERRO, TESTE ou SUCESSO")

	flag.Parse()

	fmt.Println("Parametros " + *path)
	fmt.Println("Parametros " + *logType)

	antes := time.Now()
	fmt.Println(antes.Format("2006.01.02 15:04:05"))

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	qtd := 0

	for {
		s, err := r.ReadString('\n')

		if err != nil {
			break
		}

		if strings.Contains(s, *logType) {
			qtd++
		}
	}

	fmt.Println(qtd)

	defer f.Close()

	depois := time.Now()
	fmt.Println(depois.Format("2006.01.02 15:04:05"))
}
