package rich

type TipoMovimento struct {
	CodTipoMovimento int
	Descricao string
	Fator int
	NumeroSequencia int
}

func NewTipoMovimento() ( *TipoMovimento ) {
	return &TipoMovimento{};
}