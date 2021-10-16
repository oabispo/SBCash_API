package richData

import (
	"database/sql"
	"errors"
	"github.com/oabispo/BAMySQLHelper"
	repo "sbcash_api/repositories/rich"
	"strconv"
	"strings"
)

type mySQLCupom struct {
	*repo.Cupom
}

func (p *mySQLCupom) MapFields(fetch *sql.Rows) error {
	var parametro, valor string
	err := fetch.Scan(&parametro, &valor)
	if err == nil {
		if len(strings.TrimSpace(valor)) > 0 {
			switch parametro {
			case "CUIDADO_1":
				{
					p.LinhaA = valor
				}
			case "CUIDADO_2":
				{
					p.LinhaB = valor
				}
			case "OPERACAO":
				{
					p.Operacao = repo.Ticket
					if op, err := strconv.Atoi(valor); err == nil {
						if op != 0 {
							p.Operacao = repo.Cartao
						}
					}
				}
			case "HORAS_VALIDADE":
				{
					p.DiasVigencia = 0
					if dv, err := strconv.Atoi(valor); err == nil {
						p.DiasVigencia = dv / 24
					}
				}
			case "VAL_MAX_VENDA_TICKET":
				{
					p.ValorMaxVenda = 0
					if mv, err := strconv.ParseFloat(valor, 64); err == nil {
						p.ValorMaxVenda = mv
					}

				}
			}
		}
	}

	return err
}

func Cupom_GetInfo(db *sql.DB) (*repo.Cupom, error) {
	var sb strings.Builder
	sb.WriteString("select parametro, valor from parametros where parametro in ")
	sb.WriteString("('CUIDADO_1', 'CUIDADO_2', 'OPERACAO', 'HORAS_VALIDADE', 'VAL_MAX_VENDA_TICKET')")
	var stmt string = sb.String()

	dbh := BAMySQLHelper.New(db)
	var e *mySQLCupom = &mySQLCupom{&repo.Cupom{}}
	fakeNewCupom := func() interface{} { return e }
	data, err := dbh.FetchMany(fakeNewCupom, stmt)

	if err != nil {
		return nil, err
	} else {
		if len(data) > 0 {
			var Cupom *mySQLCupom = data[0].(*mySQLCupom)
			return Cupom.Cupom, nil
		} else {
			return nil, nil
		}
	}
}

func Cupom_UpdateInfo(db *sql.DB, LinhaA string, LinhaB string, DiasVigencia int, Operacao repo.TipoOperacao, ValorMaxVenda float64) error {
	var msgsErr strings.Builder

	dbh := BAMySQLHelper.New(db)
	doInsert := func(parametro string, valor string) (interface{}, error) {
		return dbh.Insert("insert into parametros(parametro, valor) values (?, ?)", parametro, valor)
	}
	doUpdate := func(parametro string, valor string) (int64, error) {
		return dbh.Update("update parametros set valor = ? where parametro  = ?", valor, parametro)
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "CUIDADO_1"); err == nil {
		if data > 0 {
			_, err = doUpdate("CUIDADO_1", LinhaA)
		} else {
			_, err = doInsert("CUIDADO_1", LinhaA)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "CUIDADO_2"); err == nil {
		if data > 0 {
			_, err = doUpdate("CUIDADO_2", LinhaB)
		} else {
			_, err = doInsert("CUIDADO_2", LinhaB)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "OPERACAO"); err == nil {
		var op = int(Operacao)
		if data > 0 {
			_, err = doUpdate("OPERACAO", strconv.Itoa(op))
		} else {
			_, err = doInsert("OPERACAO", strconv.Itoa(op))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "HORAS_VALIDADE"); err == nil {
		if data > 0 {
			_, err = doUpdate("HORAS_VALIDADE", strconv.Itoa(DiasVigencia*24))
		} else {
			_, err = doInsert("HORAS_VALIDADE", strconv.Itoa(DiasVigencia*24))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VAL_MAX_VENDA_TICKET"); err == nil {
		if data > 0 {
			_, err = doUpdate("VAL_MAX_VENDA_TICKET", strconv.FormatFloat(ValorMaxVenda, 'f', 2, 32))
		} else {
			_, err = doInsert("VAL_MAX_VENDA_TICKET", strconv.FormatFloat(ValorMaxVenda, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if len(msgsErr.String()) > 0 {
		var mainErr error = errors.New(msgsErr.String())
		return mainErr
	} else {
		return nil
	}
}
