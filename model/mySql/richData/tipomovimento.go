package richData

import (
	"database/sql"
	"errors"
	"github.com/oabispo/BAMySQLHelper"
	repo "sbcash_api/repositories/rich"
)

type mySQLTipoMovimento struct {
	*repo.TipoMovimento
}

func newMySQLTipoMovimento() interface{} {
	return &mySQLTipoMovimento{&repo.TipoMovimento{}}
}

func (p *mySQLTipoMovimento) MapFields(fetch *sql.Rows) error {
	var NumeroSequencia sql.NullInt64
	err := fetch.Scan(&p.CodTipoMovimento, &p.Fator, &NumeroSequencia, &p.Descricao)
	if NumeroSequencia.Valid {
		p.NumeroSequencia = int(NumeroSequencia.Int64)
	}
	return err
}

func TipoMovimento_GetByID(db *sql.DB, cod_idioma int, cod_tipomovimento int) (*repo.TipoMovimento, error) {
	var stmt string = "select tm.*, tmd.descricao from tipo_movimento tm inner join tipo_mov_descr tmd on tmd.cod_tpmov = tm.cod_tpmov and tmd.cod_idioma = ? where tm.cod_tpmov = ? "

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchOne(newMySQLTipoMovimento, stmt, cod_idioma, cod_tipomovimento)

	if err == nil {
		if data != nil {
			var item *mySQLTipoMovimento = data.(*mySQLTipoMovimento)
			return item.TipoMovimento, err
		} else {
			return nil, errors.New("tipo movimento não encontrado!")
		}
	} else {
		return nil, err
	}
}

func convertRawTipoMovimento(data []interface{}, err error) ([]*repo.TipoMovimento, error) {
	if err != nil {
		return nil, err
	} else {
		var items []*repo.TipoMovimento = make([]*repo.TipoMovimento, 0, len(data))
		for _, item := range data {
			items = append(items, item.(*mySQLTipoMovimento).TipoMovimento)
		}
		return items, err
	}
}

func TipoMovimento_GetAll(db *sql.DB, cod_idioma int) ([]*repo.TipoMovimento, error) {
	var stmt string = "select tm.*, tmd.descricao from tipo_movimento tm inner join tipo_mov_descr tmd on tmd.cod_tpmov = tm.cod_tpmov and tmd.cod_idioma = ?"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchMany(newMySQLTipoMovimento, stmt, cod_idioma)
	items, err := convertRawTipoMovimento(data, err)
	return items, err
}

func TipoMovimento_Inserir(db *sql.DB, cod_idioma int, descricao string, fator int) (int64, error) {
	var stmt string = "insert into tipo_movimento (fator, numerosequencia) values(?, 0)"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.Insert(stmt, fator)
	var id int64 = -1
	if data != nil {
		id = data.(int64)
		_, err = dbh.Insert("insert into tipo_mov_descr(cod_tpmov, cod_idioma, descricao) values (?, ?, ?)", id, cod_idioma, descricao)
	}

	return id, err
}

func TipoMovimento_Atualizar(db *sql.DB, descricao string, fator int, numerosequencia int, cod_tipomovimento int, cod_idioma int) (int64, error) {
	var stmt string = "update tipo_movimento set numerosequencia = ?, fator = ? where cod_tpmov = ?"

	dbh := BAMySQLHelper.New(db)
	if result, err := dbh.Update(stmt, numerosequencia, fator, cod_tipomovimento); err == nil {
		result, err = dbh.Update("update tipo_mov_descr set descricao = ? where cod_idioma = ? and cod_tpmov = ?", descricao, cod_idioma, cod_tipomovimento)
		return result, err
	} else {
		return -1, err
	}
}

func TipoMovimento_Remover(db *sql.DB, cod_idioma int, cod_tipomovimento int) (int64, error) {
	var stmt string = "delete from tipo_mov_descr where cod_idioma = ? and cod_tpmov = ?"

	dbh := BAMySQLHelper.New(db)
	if result, err := dbh.Delete(stmt, cod_idioma, cod_tipomovimento); err == nil {
		result, err = dbh.Delete("delete from tipo_movimento where cod_tpmov = ?", cod_tipomovimento)
		return result, err
	} else {
		return -1, err
	}
}
