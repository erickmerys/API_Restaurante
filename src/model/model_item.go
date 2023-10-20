package model

type Item struct {
	ID             uint64  `json:"id"`
	NumItem        uint64  `json:"numItem"`
	ValorItem      float64 `json:"valorItem"`
	QuantidadeItem uint64  `json:"quantidadeItem"`
}
