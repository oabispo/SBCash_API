package plainData

import (
	"database/sql"
	"github.com/oabispo/BAMySQLHelper"
	repo "sbcash_api/repositories/plain"
)

type mysqlUsuario struct {
	*repo.Usuario
}

func newMySQLUsuario() interface{} {
	return &mysqlUsuario{&repo.Usuario{}}
}

func (p *mysqlUsuario) MapFields(fetch *sql.Rows) error {
	err := fetch.Scan(&p.Cod_user, &p.Cod_perfil, &p.Nome, &p.Senha, &p.Status, &p.Aut_cortesia, &p.Vis_pos_cx)
	return err
}

func Usuario_GetByID(db *sql.DB, cod_user int) (*repo.Usuario, error) {
	var stmt string = "select u.cod_user, u.cod_perfil, u.nome, u.senha, u.status, u.aut_cortesia, u.vis_pos_cx from usuario u where u.cod_user = ?"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchOne(newMySQLUsuario, stmt, cod_user)

	if err != nil {
		return nil, err
	} else {
		if data != nil {
			var usuario *mysqlUsuario = data.(*mysqlUsuario)
			return usuario.Usuario, err
		} else {
			return nil, err
		}
	}
}

func convertRawUsuario(data []interface{}, err error) ([]*repo.Usuario, error) {
	if err != nil {
		return nil, err
	} else {
		var usuarios []*repo.Usuario = make([]*repo.Usuario, 0, len(data))
		for _, item := range data {
			usuarios = append(usuarios, item.(*mysqlUsuario).Usuario)
		}
		return usuarios, err
	}
}

func Usuario_GetAll(db *sql.DB) ([]*repo.Usuario, error) {
	var stmt string = "select u.cod_user, u.cod_perfil, u.nome, u.senha, u.status, u.aut_cortesia, u.vis_pos_cx from usuario u"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchMany(newMySQLUsuario, stmt)
	usuarios, err := convertRawUsuario(data, err)
	return usuarios, err
}

func Usuario_GetAllPaged(db *sql.DB, rowsPerPage int, currentPage int) ([]*repo.Usuario, error) {
	var stmt string = "select u.cod_user, u.cod_perfil, u.nome, u.senha, u.status, u.aut_cortesia, u.vis_pos_cx from usuario u"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.FetchPage(rowsPerPage, currentPage, newMySQLUsuario, stmt)
	usuarios, err := convertRawUsuario(data, err)
	return usuarios, err
}

func Usuario_Inserir(db *sql.DB, Nome string, Senha string, Perfil int, PermiteCortesia bool, PermiteVerCaixa bool, Status int) (int64, error) {
	var stmt string = "insert into Usuario (nome, senha, cod_perfil, aut_cortesia, vis_pos_cx, status) values(?, MD5(upper(?)), ?, ?, ?, ?)"

	dbh := BAMySQLHelper.New(db)
	data, err := dbh.Insert(stmt, Nome, Senha, Perfil, PermiteCortesia, PermiteVerCaixa, Status)
	var id int64 = -1
	if data != nil {
		id = data.(int64)
	}

	return id, err
}

func Usuario_Atualizar(db *sql.DB, Cod_Usuario int, Nome string, Senha string, Perfil int, PermiteCortesia bool, PermiteVerCaixa bool, Status int) (int, error) {
	var stmt string = "update usuario set nome = ?, senha = MD5(upper(?)), cod_perfil = ?, aut_cortesia = ?, vis_pos_cx = ?, status = ? where cod_user = ?"

	dbh := BAMySQLHelper.New(db)
	rowsAffected, err := dbh.Update(stmt, Nome, Senha, Perfil, PermiteCortesia, PermiteVerCaixa, Status, Cod_Usuario)

	return int(rowsAffected), err
}

func Usuario_Remover(db *sql.DB, cod_usuario int) (int64, error) {
	var stmt string = "delete from usuario where cod_user = ?"

	dbh := BAMySQLHelper.New(db)
	deleted, err := dbh.Delete(stmt, cod_usuario)

	return deleted, err
}
