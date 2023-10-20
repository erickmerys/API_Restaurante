package controllers

import (
	"encoding/json"

	"go_modulos/src/banck"
	"go_modulos/src/model"
	"go_modulos/src/repository"
	"go_modulos/src/response"
	"io"
	"net/http"
)

func CriarPedidos(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var pedido model.Pedido
	if erro = json.Unmarshal(corpoRequest, &pedido); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banck.Conectar()
	if erro != nil {
		response.Erro(w , http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePedido(db)
	PedidoCriado, erro := repositorio.CriarPedido(pedido)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, PedidoCriado)
}
