package rich

import "time"
import "encoding/json"

type Estabelecimento struct {
	Nome string `json:"nome"`
	RazaoSocial string `json:"razaosocial"`
	CNPJ string `json:"cnpj"`
	Endereco string `json:"endereco"`
	Cidade string `json:"cidade"`
	Data_Instalacao time.Time `json:"datainstalacao"`
	Hora_Abre time.Time `json:"horaabre"`
	Hora_Fecha time.Time `json:"horafecha"`
	Tempo_login_Expira int `json:"tempologinexpira"`
	NumViasDescricao string `json:"numviasdescricao"`
	Impressao_parcial int `json:"impressaoparcial"`
	Idioma int `json:"idioma"`
}

func NewEstabelecimento() ( *Estabelecimento ) {
	return &Estabelecimento{}
}

func (e *Estabelecimento) MarshalJSON() ([]byte, error) {
	type Alias Estabelecimento;
	aux := &struct { 
		Hora_Abre string `json:"horaabre"`
		Hora_Fecha string `json:"horafecha"`
		Data_Instalacao string `json:"datainstalacao"`
		*Alias 
	}{ Hora_Abre: e.Hora_Abre.Format( "15:04:05" ), Hora_Fecha: e.Hora_Fecha.Format( "15:04:05" ), Data_Instalacao: e.Data_Instalacao.Format( "02/01/2006" ), Alias: (*Alias)(e) };
	return json.Marshal( aux );
}