package richData

import (
	"database/sql"
	"errors"
	"github.com/oabispo/BAMySQLHelper"
	repo "sbcash_api/repositories/rich"
	"strconv"
	"strings"
	"time"
)

type mySQLEstabelecimento struct {
	*repo.Estabelecimento
}

func (p *mySQLEstabelecimento) MapFields(fetch *sql.Rows) error {
	var parametro, valor string
	err := fetch.Scan(&parametro, &valor)
	if err == nil {
		if len(strings.TrimSpace(valor)) > 0 {
			switch parametro {
			case "NOME_BINGO":
				p.Nome = valor
			case "RZ_SOCIAL":
				p.RazaoSocial = valor
			case "CNPJ":
				p.CNPJ = valor
			case "ENDERECO_BINGO":
				p.Endereco = valor
			case "CIDADE":
				p.Cidade = valor
			case "DATA_INSTALACAO":
				{
					data, err := time.Parse("02/1/2006", valor)
					if err == nil {
						p.Data_Instalacao = data
					}
				}
			case "HR_ABRE":
				{
					if data, err := time.Parse("15:04:05", valor); err == nil {
						p.Hora_Abre = data
					}
					p.Hora_Abre.Format("15:04:05")
				}
			case "HR_FECHA":
				{
					if data, err := time.Parse("15:04:05", valor); err == nil {
						p.Hora_Fecha = data
					}
					p.Hora_Fecha.Format("15:04:05")
				}
			case "TEMPO_LOGIN_EXPIRA":
				{
					data, err := strconv.Atoi(valor)
					if err == nil {
						p.Tempo_login_Expira = data
					}
				}
			case "NUM_VIAS_DESCR":
				{
					p.NumViasDescricao = valor
				}
			case "IMPRESSAO_PARCIAL":
				{
					data, err := strconv.Atoi(valor)
					if err == nil {
						if data == 1 {
							p.Impressao_parcial = 1
						} else {
							p.Impressao_parcial = 0
						}
					}
				}
			case "IDIOMA":
				{
					data, err := strconv.Atoi(valor)
					if err == nil {
						p.Idioma = data
					}
				}
			}
		}
	}

	return err
}

func Estabelecimento_GetInfo(db *sql.DB) (*repo.Estabelecimento, error) {
	var sb strings.Builder
	sb.WriteString("select parametro, valor from parametros where parametro in ")
	sb.WriteString("('NOME_BINGO', 'ENDERECO_BINGO', 'CIDADE', 'RZ_SOCIAL', 'CNPJ', 'NUM_VIAS_DESCR', 'TEMPO_LOGIN_EXPIRA', 'IDIOMA', 'HR_ABRE', 'HR_FECHA', 'DATA_INSTALACAO', 'IMPRESSAO_PARCIAL' )")
	var stmt string = sb.String()

	dbh := BAMySQLHelper.New(db)
	var e *mySQLEstabelecimento = &mySQLEstabelecimento{&repo.Estabelecimento{}}
	fakeNewEstabelecimento := func() interface{} { return e }
	data, err := dbh.FetchMany(fakeNewEstabelecimento, stmt)

	if err != nil {
		return nil, err
	} else {
		if len(data) > 0 {
			var estabelecimento *mySQLEstabelecimento = data[0].(*mySQLEstabelecimento)
			return estabelecimento.Estabelecimento, nil
		} else {
			return nil, nil
		}
	}
}

func Estabelecimento_UpdateInfo(db *sql.DB, Nome string, RazaoSocial string, CNPJ string, Endereco string, Cidade string, DataInstalacao time.Time, HoraAbre time.Time, HoraFecha time.Time, TempoLoginExpira int, NumViasDescricao string, ImpressaoParcial int, Idioma int) error {
	var msgsErr strings.Builder

	dbh := BAMySQLHelper.New(db)
	doInsert := func(parametro string, valor string) (interface{}, error) {
		return dbh.Insert("insert into parametros(parametro, valor) values (?, ?)", parametro, valor)
	}

	doUpdate := func(parametro string, valor string) (int64, error) {
		return dbh.Update("update parametros set valor = ? where parametro  = ?", valor, parametro)
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "NOME_BINGO"); err == nil {
		if data > 0 {
			_, err = doUpdate("NOME_BINGO", Nome)
		} else {
			_, err = doInsert("NOME_BINGO", Nome)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "ENDERECO_BINGO"); err == nil {
		if data > 0 {
			_, err = doUpdate("ENDERECO_BINGO", Endereco)
		} else {
			_, err = doInsert("ENDERECO_BINGO", Endereco)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "CIDADE"); err == nil {
		if data > 0 {
			_, err = doUpdate("CIDADE", Cidade)
		} else {
			_, err = doInsert("CIDADE", Cidade)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "RZ_SOCIAL"); err == nil {
		if data > 0 {
			_, err = doUpdate("RZ_SOCIAL", RazaoSocial)
		} else {
			_, err = doInsert("RZ_SOCIAL", RazaoSocial)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "CNPJ"); err == nil {
		if data > 0 {
			_, err = doUpdate("CNPJ", CNPJ)
		} else {
			_, err = doInsert("CNPJ", CNPJ)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "DATA_INSTALACAO"); err == nil {
		if data > 0 {
			_, err = doUpdate("DATA_INSTALACAO", DataInstalacao.Format("02/01/2006"))
		} else {
			_, err = doInsert("DATA_INSTALACAO", DataInstalacao.Format("02/01/2006"))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "HR_ABRE"); err == nil {
		if data > 0 {
			_, err = doUpdate("HR_ABRE", HoraAbre.Format("15:04:05 "))
		} else {
			_, err = doInsert("HR_ABRE", HoraAbre.Format("15:04:05 "))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "HR_FECHA"); err == nil {
		if data > 0 {
			_, err = doUpdate("HR_FECHA", HoraFecha.Format("15:04:05 "))
		} else {
			_, err = doInsert("HR_FECHA", HoraFecha.Format("15:04:05 "))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "NUM_VIAS_DESCR"); err == nil {
		if data > 0 {
			_, err = doUpdate("NUM_VIAS_DESCR", NumViasDescricao)
		} else {
			_, err = doInsert("NUM_VIAS_DESCR", NumViasDescricao)
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "TEMPO_LOGIN_EXPIRA"); err == nil {
		if data > 0 {
			_, err = doUpdate("TEMPO_LOGIN_EXPIRA", strconv.Itoa(TempoLoginExpira))
		} else {
			_, err = doInsert("TEMPO_LOGIN_EXPIRA", strconv.Itoa(TempoLoginExpira))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "IMPRESSAO_PARCIAL"); err == nil {
		if data > 0 {
			_, err = doUpdate("IMPRESSAO_PARCIAL", strconv.Itoa(ImpressaoParcial))
		} else {
			_, err = doInsert("IMPRESSAO_PARCIAL", strconv.Itoa(ImpressaoParcial))
		}
		if err != nil {
			msgsErr.WriteString("[" + err.Error() + "]")
		}
	}

	if data, err := BAMySQLHelper.GetIntValue(db, "select count(1) from parametros where parametro = ?", "IDIOMA"); err == nil {
		if data > 0 {
			_, err = doUpdate("IDIOMA", strconv.Itoa(Idioma))
		} else {
			_, err = doInsert("IDIOMA", strconv.Itoa(Idioma))
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
