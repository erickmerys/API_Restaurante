package repository

import (
	"database/sql"
	"go_modulos/src/model"
)

type Pedido struct {
	db	*sql.DB
}

func NovoRepositorioDePedido(db *sql.DB) *Pedido {
	return &Pedido{db}
}

//CriarPedido cria um novo pedido que contém itens dentro dele
func (repositorio Pedido) CriarPedido(pedido model.Pedido) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into Pedido(numeroComanda, valorTotal) values(?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(pedido.NumeroComanda, pedido.ValoTotal)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}