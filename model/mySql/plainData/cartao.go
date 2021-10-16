package plainData

import (
	"database/sql"
	"github.com/oabispo/BAMySQLHelper"
	repo "sbcash_api/repositories/plain"
	"time"
)

type mySQLCartao struct {
	*repo.Cartao
}

func newMySQLCartao() interface{} {
	return &mySQLCartao{&repo.Cartao{}}
}

// remover assim que o Go for atualizado. Na versão Windows 10 não preciso disso.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (time.Time, error) {
	if !nt.Valid {
		return time.Time{}, nil
	}
	return nt.Time, nil
}

func (p *mySQLCartao) MapFields(fetch *sql.Rows) error {
	var Cod_cartao sql.NullInt64
	var Saldo_creditado sql.NullFloat64
	var Pin sql.NullString
	var Saldo sql.NullFloat64
	var Dh_criacao NullTime
	var Dh_encerramento NullTime
	var Pontos sql.NullInt64
	var Cod_maq_jogo sql.NullString
	var Via sql.NullInt64
	var Banda_magnetica sql.NullString
	var Cod_cliente sql.NullInt64
	var Cortesia sql.NullInt64

	err := fetch.Scan(&Cod_cartao, &Cod_cliente, &Cortesia, &Saldo_creditado, &Pin, &Saldo, &Dh_criacao, &Dh_encerramento, &Pontos, &Cod_maq_jogo, &Via, &Banda_magnetica)

	p.Cod_cartao = int(Cod_cartao.Int64)
	p.Cod_cliente = int(Cod_cliente.Int64)
	p.Cortesia = int(Cortesia.Int64)
	p.Saldo_creditado = Saldo_creditado.Float64
	p.Pin = Pin.String
	p.Saldo = Saldo.Float64

	if Dh_criacao.Valid {
		p.Dh_criacao = Dh_criacao.Time
	}

	if Dh_encerramento.Valid {
		p.Dh_encerramento = Dh_encerramento.Time
	}

	p.Pontos = int(Pontos.Int64)

	if Cod_maq_jogo.Valid {
		p.Cod_maq_jogo = Cod_maq_jogo.String
	}

	if Via.Valid {
		p.Via = int(Via.Int64)
	}

	if Banda_magnetica.Valid {
		p.Banda_magnetica = Banda_magnetica.String
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func Cartao_GetByID(db *sql.DB, cod_cartao int) (*repo.Cartao, error) {
	var stmt string = "select c.Cod_cartao, c.Cod_cliente, c.Cortesia, c.Saldo_creditado, c.Pin, c.Saldo, c.Dh_criacao, c.Dh_encerramento, c.Pontos, c.Cod_maq_jogo, c.Via, c.banda_magnetica from cartao c where c.cod_cartao = ?"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchOne(newMySQLCartao, stmt, cod_cartao)

	if err != nil {
		return nil, err
	} else {
		var cartao *repo.Cartao = (data.(*mySQLCartao)).Cartao
		return cartao, err
	}
}
