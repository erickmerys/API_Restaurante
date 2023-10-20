package main

import (
	"fmt"
	"go_modulos/src/config"
	rotaconfiguracao "go_modulos/src/router/rota_configuracao"

	"log"
	"net/http"
)

func main() {
	config.Carregar()

	r := rotaconfiguracao.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
