package endpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	router "sbcash_api/BARoutingHelper"

	usuario_sql "sbcash_api/model/mySql/plainData"

	repo "sbcash_api/repositories/plain"
)

func NewRouterUsuario(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	usuario := router.NewBAUrlSegmentRoute(router.MT_GET, "/usuario", segmentParent)

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		router.ResponseTextAsHTML(w, http.StatusNotFound, "<h1>Nenhum usuario encontrado!</h1>")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		items, err := usuario_sql.Usuario_GetAll(db)
		if err == nil {
			if len(items) < 1 {
				handler_nf(w, r)
			} else {
				GETResponse(w, items)
			}
		} else {
			ERRORResponse(w, err)
		}
	}
	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	usuario.AddSubRouter(NewRouterUsuarioPorId(db, usuario))
	usuario.AddSubRouter(NewRouterUsuario_Post(db, usuario))
	usuario.AddSubRouter(NewRouterUsuario_Put(db, usuario))
	usuario.AddSubRouter(NewRouterUsuario_Delete(db, usuario))

	return usuario
}

func NewRouterUsuarioPorId(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	usuario := router.NewBAUrlSegmentRoute(router.MT_GET, "/([0-9]+)", segmentParent)

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_usuario, _ := strconv.Atoi(basm.LastSegment())

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("<h1>Usuario com cod_usuario = %v não encontrado!</h1>", cod_usuario)))
		fmt.Fprint(w)
	}

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_usuario, _ := strconv.Atoi(basm.LastSegment())

		items, err := usuario_sql.Usuario_GetByID(db, cod_usuario)
		if err == nil {
			if items == nil {
				handler_nf(w, r)
			} else {
				GETResponse(w, items)
			}
		} else {
			ERRORResponse(w, err)
		}
	}

	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return usuario
}

func NewRouterUsuario_Post(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	usuario := router.NewBAUrlSegmentRoute(router.MT_POST, "", segmentParent)
	var cod_usuario int64
	var err error

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_usuario))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		u := repo.NewUsuario()
		d := json.NewDecoder(r.Body)
		d.Decode(&u)

		cod_usuario, err = usuario_sql.Usuario_Inserir(db, u.Nome, u.Senha, u.Cod_perfil, u.Aut_cortesia, u.Vis_pos_cx, u.Status)
		if err == nil {
			if cod_usuario > 0 {
				strId := strconv.Itoa(int(cod_usuario))
				POSTResponse(w, map[string]string{"id": strId})
			} else {
				handler_nf(w, r)
			}
		} else {
			data := map[string]string{"id": strconv.Itoa(int(cod_usuario))}
			data["message"] = err.Error()
			ERRORResponse(w, data)
		}
	}

	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return usuario
}

func NewRouterUsuario_Put(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	usuario := router.NewBAUrlSegmentRoute(router.MT_PUT, "/([0-9]+)", segmentParent)
	var cod_usuario int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_usuario))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_usuario, _ = strconv.Atoi(basm.LastSegment())

		u := repo.NewUsuario()
		d := json.NewDecoder(r.Body)
		d.Decode(&u)

		rowsAffected, err := usuario_sql.Usuario_Atualizar(db, cod_usuario, u.Nome, u.Senha, u.Cod_perfil, u.Aut_cortesia, u.Vis_pos_cx, u.Status)
		if err == nil {
			if rowsAffected > 0 {
				strId := strconv.Itoa(int(cod_usuario))
				PUTResponse(w, map[string]string{"id": strId})
			} else {
				handler_nf(w, r)
			}
		} else {
			data := map[string]string{"id": strconv.Itoa(cod_usuario)}
			data["message"] = err.Error()
			ERRORResponse(w, data)
		}
	}

	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return usuario
}

func NewRouterUsuario_Delete(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	usuario := router.NewBAUrlSegmentRoute(router.MT_DELETE, "/([0-9]+)", segmentParent)
	var cod_usuario int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_usuario))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_usuario, _ = strconv.Atoi(basm.LastSegment())

		u := repo.NewUsuario()
		d := json.NewDecoder(r.Body)
		d.Decode(&u)

		rowsAffected, err := usuario_sql.Usuario_Remover(db, cod_usuario)
		if err == nil {
			if rowsAffected > 0 {
				strId := strconv.Itoa(int(cod_usuario))
				DELETEResponse(w, map[string]string{"id": strId})
			} else {
				handler_nf(w, r)
			}
		} else {
			data := map[string]string{"id": strconv.Itoa(cod_usuario)}
			data["message"] = err.Error()
			ERRORResponse(w, data)
		}
	}

	usuario.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return usuario
}
