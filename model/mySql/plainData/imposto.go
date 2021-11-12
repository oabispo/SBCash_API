package plainData

import (
	"database/sql"
	"errors"
	repo "sbcash_api/repositories/plain"

	"github.com/oabispo/BAMySQLHelper"
)

type mysqlImposto struct {
	*repo.Imposto
}

func newMySQLImposto() interface{} {
	return &mysqlImposto{&repo.Imposto{}}
}

func (p *mysqlImposto) MapFields(fetch *sql.Rows) error {
	err := fetch.Scan(&p.Cod_imposto, &p.Descricao, &p.Aliquota)
	return err
}

func Imposto_GetByID(db *sql.DB, cod_imposto int) (*repo.Imposto, error) {
	var stmt string = "select i.cod_imposto, i.descricao, i.aliquota from Imposto i where i.cod_imposto = ?"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchOne(newMySQLImposto, stmt, cod_imposto)

	if err == nil {
		if data != nil {
			var imposto *mysqlImposto = data.(*mysqlImposto)
			return imposto.Imposto, err
		} else {
			return nil, errors.New("Imposto não encontrado!")
		}
	} else {
		return nil, err
	}
}

func convertRawImposto(data []interface{}, err error) ([]*repo.Imposto, error) {
	if err != nil {
		return nil, err
	} else {
		var impostos []*repo.Imposto = make([]*repo.Imposto, 0, len(data))
		for _, item := range data {
			sqlImposto := item.(*mysqlImposto)
			impostos = append(impostos, sqlImposto.Imposto)
		}
		return impostos, err
	}
}

func Imposto_GetAll2(db *sql.DB) ([]*repo.Imposto, error) {
	var stmt string = "select i.cod_imposto, i.descricao, i.aliquota from Imposto i"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchMany(newMySQLImposto, stmt)
	impostos, err := convertRawImposto(data, err)
	return impostos, err
}

func Imposto_GetAll(db *sql.DB) ([]*repo.Imposto, error) {
	var stmt string = "select i.cod_imposto, i.descricao, i.aliquota from Imposto i"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchMany(newMySQLImposto, stmt)

	createItemCallback := func(item interface{}) interface{} {
		sqlImposto := item.(*mysqlImposto)
		imposto := &repo.Imposto{}
		imposto = sqlImposto.Imposto
		return imposto
	}

	items := make([]interface{}, 0, len(data))
	ConvertRawData(data, items, createItemCallback)
	//impostos := items.([]repo.Imposto)
	return nil, err
}

func Imposto_Inserir(db *sql.DB, descricao string, aliquota float64) (int64, error) {
	var stmt string = "insert into imposto (descricao, aliquota) values(?, ?)"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.Insert(stmt, descricao, aliquota)
	var id int64 = -1
	if data != nil {
		id = data.(int64)
	}

	return id, err
}

func Imposto_Atualizar(db *sql.DB, descricao string, aliquota float64, cod_imposto int) (int64, error) {
	var stmt string = "update imposto set descricao = ?, aliquota = ? where cod_imposto = ?"

	dbh := BAMySQLHelper.New(db)
	if result, err := dbh.Update(stmt, descricao, aliquota, cod_imposto); err == nil {
		return result, nil
	} else {
		return -1, err
	}
}

func Imposto_Remover(db *sql.DB, cod_imposto int) (int64, error) {
	var stmt string = "delete from imposto where cod_imposto = ?"

	dbh := BAMySQLHelper.New(db)
	if result, err := dbh.Delete(stmt, cod_imposto); err == nil {
		return result, nil
	} else {
		return -1, err
	}
}
