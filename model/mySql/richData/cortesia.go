package richData

import (
	"database/sql"
	"github.com/oabispo/BAMySQLHelper"

	"errors"
	repo "sbcash_api/repositories/rich"
	"strconv"
	"strings"
)

type mySQLCortesia struct {
	*repo.Cortesia
}

func (p *mySQLCortesia) MapFields(fetch *sql.Rows) error {
	var parametro, valor string
	err := fetch.Scan(&parametro, &valor)
	if err == nil {
		if len(strings.TrimSpace(valor)) > 0 {
			switch parametro {
			case "VAL_MAX_CORTESIA":
				{
					p.ValorMaximo = 0
					if vm, err := strconv.ParseFloat(valor, 64); err == nil {
						p.ValorMaximo = vm
					}
				}
			case "VALIDADE_CORTESIA":
				{
					p.Vigencia = 0
					if vi, err := strconv.Atoi(valor); err == nil {
						p.Vigencia = vi / 24
					}
				}
			case "REGRA_CORTESIA":
				{
					if valor == "NP" {
						p.RegraCortesia = repo.ND
					} else if valor == "DP" {
						p.RegraCortesia = repo.DP
					}
				}
			case "VC1":
				{
					p.Valor1 = 0
					if v, err := strconv.ParseFloat(valor, 64); err == nil {
						p.Valor1 = v
					}
				}
			case "VC2":
				{
					p.Valor2 = 0
					if v, err := strconv.ParseFloat(valor, 64); err == nil {
						p.Valor2 = v
					}
				}
			case "VC3":
				{
					p.Valor3 = 0
					if v, err := strconv.ParseFloat(valor, 64); err == nil {
						p.Valor3 = v
					}
				}
			case "VC4":
				{
					p.Valor4 = 0
					if v, err := strconv.ParseFloat(valor, 64); err == nil {
						p.Valor4 = v
					}
				}
			case "VC5":
				{
					p.Valor5 = 0
					if v, err := strconv.ParseFloat(valor, 64); err == nil {
						p.Valor5 = v
					}
				}
			}
		}
	}

	return err
}

func Cortesia_GetInfo(db *sql.DB) (*repo.Cortesia, error) {
	var sb strings.Builder
	sb.WriteString("select parametro, valor from parametros where parametro in ")
	sb.WriteString("('VAL_MAX_CORTESIA', 'VALIDADE_CORTESIA', 'REGRA_CORTESIA', 'VC1', 'VC2', 'VC3', 'VC4', 'VC5' )")
	var stmt string = sb.String()

	dbh := BAMySQLHelper.New(db)
	var e *mySQLCortesia = &mySQLCortesia{&repo.Cortesia{}}
	fakeNewCortesia := func() interface{} { return e }
	data, err := dbh.FetchMany(fakeNewCortesia, stmt)

	if err != nil {
		return nil, err
	} else {
		if len(data) > 0 {
			var Cortesia *mySQLCortesia = data[0].(*mySQLCortesia)
			return Cortesia.Cortesia, nil
		} else {
			return nil, nil
		}
	}
}

func Cortesia_Saldo(db *sql.DB) (float64, error) {
	var stmt string = "select parametro, valor from parametros where parametro = 'SALDO_CORTESIA'"
	var saldo float64 = 0
	if data, err := BAMySQLHelper.GetStringValue(db, stmt); err == nil {
		saldo, err = strconv.ParseFloat(data, 64)
		return saldo, err
	} else {
		return saldo, nil
	}
}

func Cortesia_UpdateSaldo(db *sql.DB, saldo float64) (int64, error) {
	var stmt string = "update parametros set valor = ? where parametro = 'SALDO_CORTESIA'"
	dbh := BAMySQLHelper.New(db)
	total, err := dbh.Update(stmt, saldo)
	return total, err
}

func Cortesia_UpdateInfo(db *sql.DB, ValorMaximo float64, Vigencia int, RegraCortesia repo.RegraCortesia, Valor1 float64, Valor2 float64, Valor3 float64, Valor4 float64, Valor5 float64) error {
	var msgsErr strings.Builder

	dbh := BAMySQLHelper.New(db)
	doInsert := func(parametro string, valor string) (interface{}, error) {
		return dbh.Insert("insert into parametros(parametro, valor) values (?, ?)", parametro, valor)
	}
	doUpdate := func(parametro string, valor string) (int64, error) {
		return dbh.Update("update parametros set valor = ? where parametro  = ?", valor, parametro)
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VAL_MAX_CORTESIA"); err == nil {
		if data > 0 {
			_, err = doUpdate("VAL_MAX_CORTESIA", strconv.FormatFloat(ValorMaximo, 'f', 2, 32))
		} else {
			_, err = doInsert("VAL_MAX_CORTESIA", strconv.FormatFloat(ValorMaximo, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VALIDADE_CORTESIA"); err == nil {
		if data > 0 {
			_, err = doUpdate("VALIDADE_CORTESIA", strconv.Itoa(Vigencia*24))
		} else {
			_, err = doInsert("VALIDADE_CORTESIA", strconv.Itoa(Vigencia*24))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "REGRA_CORTESIA"); err == nil {
		rc := "DP"
		if RegraCortesia != repo.DP {
			rc = "ND"
		}
		if data > 0 {
			_, err = doUpdate("REGRA_CORTESIA", rc)
		} else {
			_, err = doInsert("REGRA_CORTESIA", rc)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VC1"); err == nil {
		if data > 0 {
			_, err = doUpdate("VC1", strconv.FormatFloat(Valor1, 'f', 2, 32))
		} else {
			_, err = doInsert("VC1", strconv.FormatFloat(Valor1, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VC2"); err == nil {
		if data > 0 {
			_, err = doUpdate("VC2", strconv.FormatFloat(Valor2, 'f', 2, 32))
		} else {
			_, err = doInsert("VC2", strconv.FormatFloat(Valor2, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VC3"); err == nil {
		if data > 0 {
			_, err = doUpdate("VC3", strconv.FormatFloat(Valor3, 'f', 2, 32))
		} else {
			_, err = doInsert("VC3", strconv.FormatFloat(Valor3, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VC4"); err == nil {
		if data > 0 {
			_, err = doUpdate("VC4", strconv.FormatFloat(Valor4, 'f', 2, 32))
		} else {
			_, err = doInsert("VC4", strconv.FormatFloat(Valor4, 'f', 2, 32))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "VC5"); err == nil {
		if data > 0 {
			_, err = doUpdate("VC5", strconv.FormatFloat(Valor5, 'f', 2, 32))
		} else {
			_, err = doInsert("VC5", strconv.FormatFloat(Valor5, 'f', 2, 32))
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
