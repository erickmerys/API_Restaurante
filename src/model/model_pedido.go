package model

type Pedido struct {
	ID            uint64  `json:"id"`
	NumeroComanda uint64  `json:"numeroComanda"`
	ValoTotal     float64 `json:"valorTotal"`
}
