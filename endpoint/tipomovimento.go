package endpoint

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	router "sbcash_api/BARoutingHelper"

	TipoMovimento_sql "sbcash_api/model/mySql/richData"

	repo "sbcash_api/repositories/rich"
)

func NewRouterTipoMovimento(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/tipomovimento", segmentParent)

	var cod_tipomovimento int

	handler_err := func(w http.ResponseWriter, r *http.Request) {
		ERRORResponse(w, nil)
	}

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		NOTFOUNDResponse(w, nil)
	}

	handler_nf_byid := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_tipomovimento))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		baqm := router.NewBAUrlQueryManage(r.URL)
		var cod_idioma int
		if value, found := baqm.GetKeyValue("idioma"); found {
			cod_idioma, _ = strconv.Atoi(value)
		}

		if value, found := baqm.GetKeyValue("codtipomovimento"); found {
			cod_tipomovimento, _ = strconv.Atoi(value)
			data, err := TipoMovimento_sql.TipoMovimento_GetByID(db, cod_idioma, cod_tipomovimento)
			if err == nil {
				if data == nil {
					handler_nf_byid(w, r)
				} else {
					GETResponse(w, data)
				}
			} else {
				handler_err(w, r)
			}
		} else {
			data, err := TipoMovimento_sql.TipoMovimento_GetAll(db, cod_idioma)
			if err == nil {
				if len(data) < 1 {
					handler_nf(w, r)
				} else {
					GETResponse(w, data)
				}
			} else {
				handler_err(w, r)
			}
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	segment.AddSubRouter(NewRouterTipoMovimento_Post(db, segment))
	segment.AddSubRouter(NewRouterTipoMovimento_Update(db, segment))
	segment.AddSubRouter(NewRouterTipoMovimento_Delete(db, segment))

	return segment
}

func NewRouterTipoMovimento_Post(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_POST, "", segmentParent)

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		ERRORResponse(w, nil)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		baqm := router.NewBAUrlQueryManage(r.URL)
		var cod_idioma int
		if value, found := baqm.GetKeyValue("idioma"); found {
			cod_idioma, _ = strconv.Atoi(value)
		}

		decode := json.NewDecoder(r.Body)
		tipoMovimento := repo.NewTipoMovimento()
		decode.Decode(&tipoMovimento)

		id, err := TipoMovimento_sql.TipoMovimento_Inserir(db, cod_idioma, tipoMovimento.Descricao, tipoMovimento.Fator)
		if err == nil {
			strId := strconv.Itoa(int(id))
			POSTResponse(w, map[string]string{"id": strId})
		} else {
			handler_nf(w, r)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}

func NewRouterTipoMovimento_Delete(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_DELETE, "", segmentParent)

	var cod_tipomovimento int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_tipomovimento))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		baqm := router.NewBAUrlQueryManage(r.URL)
		var cod_idioma int
		if value, found := baqm.GetKeyValue("idioma"); found {
			cod_idioma, _ = strconv.Atoi(value)
		}

		if value, found := baqm.GetKeyValue("codtipomovimento"); found {
			cod_tipomovimento, _ = strconv.Atoi(value)
		}

		rowsAffected, err := TipoMovimento_sql.TipoMovimento_Remover(db, cod_idioma, cod_tipomovimento)
		if (err == nil) && (rowsAffected > 0) {
			var response map[string]string = make(map[string]string)
			response["id"] = strconv.Itoa(cod_tipomovimento)
			response["totalRemovido"] = strconv.Itoa(int(rowsAffected))
			DELETEResponse(w, response)
		} else {
			handler_nf(w, r)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}

func NewRouterTipoMovimento_Update(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_PUT, "", segmentParent)

	var cod_tipomovimento int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_tipomovimento))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		baqm := router.NewBAUrlQueryManage(r.URL)
		var cod_idioma int
		var cod_tipomovimento int
		if value, found := baqm.GetKeyValue("idioma"); found {
			cod_idioma, _ = strconv.Atoi(value)
		}

		if value, found := baqm.GetKeyValue("codtipomovimento"); found {
			cod_tipomovimento, _ = strconv.Atoi(value)
		}

		decode := json.NewDecoder(r.Body)
		tipoMovimento := repo.NewTipoMovimento()
		decode.Decode(&tipoMovimento)

		rowsAffected, err := TipoMovimento_sql.TipoMovimento_Atualizar(db, tipoMovimento.Descricao, tipoMovimento.Fator, tipoMovimento.NumeroSequencia, cod_tipomovimento, cod_idioma)
		if (err == nil) && (rowsAffected > 0) {
			strId := strconv.Itoa(int(cod_tipomovimento))
			PUTResponse(w, map[string]string{"id": strId})
		} else {
			handler_nf(w, r)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}
