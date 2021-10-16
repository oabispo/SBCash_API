package plain

import "time"

type Cartao struct {
	Cod_cartao int
	Cod_cliente int
	Cortesia int
	Saldo_creditado float64
	Pin string
	Saldo float64
	Dh_criacao time.Time
	Dh_encerramento time.Time
	Pontos int
	Cod_maq_jogo string
	Via int
	Banda_magnetica string
}

func NewCartao() *Cartao {
	return &Cartao{};
}
