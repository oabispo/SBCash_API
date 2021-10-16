package plain

type Imposto struct {
	Cod_imposto int
	Descricao string
	Aliquota float64
}

func NewImposto() ( *Imposto ) {
	return &Imposto{};
}