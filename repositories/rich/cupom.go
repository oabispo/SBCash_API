package rich

type TipoOperacao int

const (
	Ticket TipoOperacao = iota
	Cartao 
)

type Cupom struct {
	LinhaA string
	LinhaB string
	DiasVigencia int
	Operacao TipoOperacao
	ValorMaxVenda float64
}

func NewCupom() *Cupom  {
	return &Cupom{};
}