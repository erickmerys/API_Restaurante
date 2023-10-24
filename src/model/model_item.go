package model

type Item struct {
	ID             uint64  `json:"id"`
	NomeItem       string  `json:"nomeItem"`
	ValorItem      float64 `json:"valorItem"`
	QuantidadeItem uint64  `json:"quantidadeItem"`
	PedidoID       uint64  `json:"pedido_id"`
}
