package rich

type RegraCortesia int

const (
	ND RegraCortesia = iota
	DP
)

type Cortesia struct {
	ValorMaximo float64
	Vigencia int
	RegraCortesia RegraCortesia
	Valor1 float64
	Valor2 float64
	Valor3 float64
	Valor4 float64
	Valor5 float64
}

func NewCortesia() *Cortesia  {
	return &Cortesia{}
}