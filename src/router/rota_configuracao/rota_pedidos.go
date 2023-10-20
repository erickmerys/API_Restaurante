package rotaconfiguracao

import (
	"go_modulos/src/controllers"

	"github.com/gorilla/mux"
)

func Gerar () *mux.Router {
	r  := mux.NewRouter()

	r.HandleFunc("/criar-pedido", controllers.CriarPedidos).Methods("POST")

	return r
}