package repository

import (
	"database/sql"
	"go_modulos/src/model"
)

type Item struct {
	db *sql.DB
}

func NovoRepositorioDeItem(db *sql.DB) *Item {
	return &Item{db}
}

func (repositorio Item) CriarItem(Item model.Item) (uint64, error){
	statement, erro := repositorio.db.Prepare("insert into Item(nomeItem, valorItem, quantidade, pedido_id) values(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(Item.NomeItem, Item.ValorItem, Item.QuantidadeItem, Item.PedidoID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), erro
}