package controllers

import (
	"encoding/json"
	"go_modulos/src/bank"
	"go_modulos/src/model"
	"go_modulos/src/repository"
	"go_modulos/src/response"
	"io"
	"net/http"
)

func CriarItem(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var item model.Item
	if erro = json.Unmarshal(corpoRequest, &item); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := bank.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repository.NovoRepositorioDeItem(db)
	ItemCriado, erro := repositorio.CriarItem(item)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, ItemCriado)
}
